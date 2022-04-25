// Copyright (c) Facebook, Inc. and its affiliates.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the "hack" directory of this source tree.

use anyhow::Result;
use cbindgen::{Builder, Config, EnumConfig, Language, MacroExpansionConfig, RenameRule};
use clap::Parser;
use std::path::PathBuf;

#[derive(Debug, Parser)]
#[clap(
    name = "ffi_cbindgen",
    about = r#"
Generate a cbindgen style C++ header for a list of .rs sources.

Example invocation:

  $ ffi_cbindgen --header foo-bar.h \
  --namespaces HPHP,hackc \
  --includes hphp/hack/src/baz/baz.h \
  hphp/hack/src/foo/foo.rs \
  hphp/hack/src/bar/bar.rs
"#
)]
struct Opt {
    /// Input files
    srcs: Vec<PathBuf>,

    /// The header file to write
    #[clap(long = "header")]
    header: PathBuf,

    /// Any namespaces to wrap the generated code in
    #[clap(long = "namespaces", use_value_delimiter = true)]
    namespaces: Vec<String>,

    /// Any headers to add to the top of the header
    #[clap(long = "includes", use_value_delimiter = true)]
    includes: Vec<String>,
}

fn builder(opts: &Opt) -> Builder {
    let mut builder = Builder::new()
        .with_config(Config {
            language: Language::Cxx,
            macro_expansion: MacroExpansionConfig { bitflags: true },
            enumeration: EnumConfig {
                rename_variant_name_fields: RenameRule::None,
                ..Default::default()
            },
            ..Default::default()
        })
        .with_header(
            "\
// Copyright (c) Facebook, Inc. and its affiliates.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the \"hack\" directory of this source tree.
",
        )
        .with_autogen_warning("// Warning, this file is autogenerated. Don't modify it manually!")
        .with_pragma_once(true);
    for src in opts.srcs.iter() {
        builder = builder.with_src(src)
    }
    for inc in opts.includes.iter().filter(|f| !f.is_empty()) {
        builder = builder.with_include(inc)
    }
    match &opts.namespaces[..] {
        [namespace, namespaces @ ..] => builder
            .with_namespace(namespace)
            .with_namespaces(namespaces),
        [] => builder,
    }
}

fn main() -> Result<()> {
    let opts = Opt::parse();
    let header = &opts.header;
    builder(&opts).generate()?.write_to_file(header);
    Ok(())
}
