#!/usr/bin/env rust-script
/*
 * Copyright (c) Meta Platforms, Inc. and affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

//!
//! ```cargo
//! [dependencies]
//! anyhow = "1.0.72"
//! atomicwrites = "0.4.1"
//! yaml-rust = "0.4.5"
//! ```

use std::io::Write;
use std::path::Path;

use anyhow::anyhow;
use anyhow::bail;
use atomicwrites::AllowOverwrite;
use atomicwrites::AtomicFile;
use yaml_rust::yaml::Hash;
use yaml_rust::Yaml;
use yaml_rust::YamlEmitter;
use yaml_rust::YamlLoader;

type Pattern = Box<dyn Fn(&str, &str) -> String + 'static>;

fn parse_pattern(pattern: &String) -> Option<(Pattern, Vec<&'static str>)> {
    // TODO: generalize this with an external table or something
    if pattern.contains("%UBUNTU_LTS_VERSION%") {
        Some((
            Box::new(move |p: &str, s: &str| p.replace("%UBUNTU_LTS_VERSION%", s)),
            vec!["20", "22"],
        ))
    } else if pattern.contains("%FEDORA_STABLE_VERSION%") {
        Some((
            Box::new(move |p: &str, s: &str| p.replace("%FEDORA_STABLE_VERSION%", s)),
            vec!["36"],
        ))
    } else {
        None
    }
}

fn substitute_value(pattern: &Pattern, value: &Yaml, substitution: &str) -> Yaml {
    match value {
        Yaml::String(str) => Yaml::String(pattern(&str, substitution)),
        Yaml::Array(arr) => Yaml::Array(
            arr.iter()
                .map(|v| substitute_value(pattern, v, substitution))
                .collect(),
        ),
        Yaml::Hash(hm) => {
            let mut new = Hash::new();
            for (key, value) in hm.iter() {
                new.insert(
                    substitute_value(pattern, &key, substitution),
                    substitute_value(pattern, &value, substitution),
                );
            }
            Yaml::Hash(new)
        }
        otherwise => otherwise.clone(),
    }
}

fn expand_value(value: &Yaml) -> Yaml {
    match value {
        Yaml::Array(arr) => Yaml::Array(arr.iter().map(expand_value).collect()),
        Yaml::Hash(hm) => {
            let mut new = Hash::new();
            for (key, value) in hm.iter() {
                //println!("key = {:?}", key);
                if let Yaml::String(key_str) = key {
                    if let Some((pat, substitutions)) = parse_pattern(&key_str) {
                        for substitution in substitutions {
                            new.insert(
                                Yaml::String(pat(&key_str, substitution)),
                                substitute_value(&pat, &value, substitution),
                            );
                        }
                    } else {
                        new.insert(key.clone(), expand_value(value));
                    }
                } else {
                    // What does a non-string key mean anyways?
                    new.insert(key.clone(), expand_value(value));
                }
            }
            Yaml::Hash(new)
        }
        otherwise => otherwise.clone(),
    }
}

fn main() -> anyhow::Result<()> {
    let Some(base_path) = std::env::var_os("RUST_SCRIPT_BASE_PATH") else {
        bail!("RUST_SCRIPT_BASE_PATH must be set");
    };
    let base_path = Path::new(&base_path);
    let input_path = base_path.join("release.yml.in");
    let output_path = base_path.join("release.yml");

    let contents = match std::fs::read_to_string(&input_path) {
        Ok(contents) => contents,
        Err(e) => {
            bail!("Unable to read {}: {}", input_path.display(), e);
        }
    };
    let doc = match YamlLoader::load_from_str(&contents) {
        Ok(doc) => doc,
        Err(e) => {
            bail!("Unable to parse {}: {}", input_path.display(), e);
        }
    };

    let new_doc: Vec<Yaml> = doc.iter().map(expand_value).collect();

    let mut out_str = String::new();
    out_str.extend("# \x40generated by generate-release-yml.rs\n".chars());

    let mut emitter = YamlEmitter::new(&mut out_str);
    for doc in new_doc {
        match emitter.dump(&doc) {
            Ok(()) => (),
            Err(e) => {
                bail!("Unable to generate {}: {}", output_path.display(), e);
            }
        }
    }

    let af = AtomicFile::new(&output_path, AllowOverwrite);
    af.write(|f| f.write_all(&out_str.as_bytes()))
        .map_err(|e| anyhow!("Unable to write {}: {}", output_path.display(), e))
}
