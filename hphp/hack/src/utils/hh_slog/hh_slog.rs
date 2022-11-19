// Copyright (c) Facebook, Inc. and its affiliates.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the "hack" directory of this source tree.

use slog::o;
use slog::Drain;

const TIMESTAMP_FORMAT: &str = "[%Y-%m-%d %H:%M:%S%.3f]";

fn timestamp_format(io: &mut dyn std::io::Write) -> std::io::Result<()> {
    write!(io, "{}", chrono::Local::now().format(TIMESTAMP_FORMAT))
}

#[derive(Clone, Debug)]
pub struct Log {
    /// Logs to both scuba and the log file.
    pub scuba: slog::Logger,
    /// Only logs to the log file.
    pub file: slog::Logger,
}

/// Creates a logger that drains to the given path. Also returns its guard (read
/// `init` for more details).
/// The file logger uses unix flock to lock the log file when flushing its logs.
/// With this overhead, the throughput is still around 7200 short messages per
/// second.
pub fn init_file(filename: &std::path::Path) -> (slog::Logger, slog_async::AsyncGuard) {
    let drain = locked_file_drain::LockedFileDrain::new(filename);
    // NB: panicking in the async thread does not cause Fuse to panic.
    let (drain, guard) = slog_async::Async::new(drain)
        .thread_name("slog_logfile".to_owned())
        .build_with_guard();

    let log = slog::Logger::root(drain.fuse(), o!());
    (log, guard)
}

pub fn init_term(binary_name: &'static str) -> (slog::Logger, slog_async::AsyncGuard) {
    let decorator = slog_term::TermDecorator::new().build();

    let drain = slog_term::FullFormat::new(decorator)
        .use_custom_timestamp(timestamp_format)
        .build()
        .fuse();

    let drain = slog_envlogger::new(drain);

    let (drain, guard) = slog_async::Async::new(drain)
        .thread_name("slog_async".to_owned())
        .build_with_guard();

    let drain = drain.fuse();

    let log = if binary_name.is_empty() {
        slog::Logger::root(drain, o!())
    } else {
        slog::Logger::root(drain, o!("bin" => binary_name))
    };
    (log, guard)
}

pub fn init_term_testing() -> slog::Logger {
    // Due to how cargo and buck capture stdout, we have to use a special sink
    // to display stdout in tests.
    let decorator = slog_term::PlainSyncDecorator::new(slog_term::TestStdoutWriter);
    let drain = slog_term::FullFormat::new(decorator)
        .use_custom_timestamp(timestamp_format)
        .build()
        .fuse();

    slog::Logger::root(drain, o!())
}

/// Initializes a `Log` where each logger prints to stdout synchronously.
/// Used for testing.
pub fn init_testing() -> Log {
    Log {
        scuba: init_term_testing(),
        file: init_term_testing(),
    }
}

/// Helper class to print up to the first 'n' items of a slice or BTreeSet.
/// It delegates to the Debug or Display traits for T.
/// If it can't show all items, then it shows "... [LEN items]" at the end
/// See docs for FmtN::slice and FmtN::bset for usage.
pub struct FmtN<'a, T> {
    /// Stores the first n items
    items: Vec<&'a T>,
    /// this is LEN, the length of the original slice or BTreeSet
    len: usize,
}

impl<'a, T> FmtN<'a, T> {
    /// Prints up to first 'n' items of the slice, e.g.
    /// println!("strings: {}", FmtN::slice(3, &names));
    /// println!("pathbufs: {:?}", FmtN::slice(5, &files));
    pub fn slice(n: usize, slice: &'a [T]) -> Self {
        Self {
            items: slice.iter().take(n).collect(),
            len: slice.len(),
        }
    }

    /// Prints up to the first 'n' items of the BTreeSet in iter() order, e.g.
    /// println!("strings: {}", FmtN::bset(3, &names));
    /// println!("pathbufs: {:?}", FmtN::bset(5, &files));
    pub fn bset(n: usize, set: &'a std::collections::BTreeSet<T>) -> Self {
        Self {
            items: set.iter().take(n).collect(),
            len: set.len(),
        }
    }
}

impl<'a, T: std::fmt::Display> std::fmt::Display for FmtN<'a, T> {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let mut sep = "";
        for item in self.items.iter() {
            f.write_str(sep)?;
            item.fmt(f)?; // Display::fmt
            sep = ", ";
        }
        if self.items.is_empty() {
            f.write_str("[none]")?;
        } else if self.len > self.items.len() {
            f.write_fmt(std::format_args!(", ... [{} total]", self.len))?;
        }
        Ok(())
    }
}

impl<'a, T: std::fmt::Debug> std::fmt::Debug for FmtN<'a, T> {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let mut sep = "";
        for item in self.items.iter() {
            f.write_str(sep)?;
            item.fmt(f)?; // Display::fmt
            sep = ", ";
        }
        if self.items.is_empty() {
            f.write_str("[none]")?;
        } else if self.len > self.items.len() {
            f.write_fmt(std::format_args!(", ... [{} total]", self.len))?;
        }
        Ok(())
    }
}

#[test]
fn test_fmtn() {
    // Test FmtN::slice
    let v0: Vec<&str> = vec![];
    let v1: Vec<&str> = vec!["a"];
    let v2: Vec<&str> = vec!["a", "b"];
    let v3: Vec<&str> = vec!["a", "b", "c"];
    assert_eq!(format!("{}", FmtN::slice(0, &v0)), "[none]".to_owned());
    assert_eq!(format!("{}", FmtN::slice(1, &v0)), "[none]".to_owned());
    assert_eq!(format!("{}", FmtN::slice(0, &v1)), "[none]".to_owned());
    assert_eq!(format!("{}", FmtN::slice(1, &v1)), "a".to_owned());
    assert_eq!(format!("{}", FmtN::slice(2, &v1)), "a".to_owned());
    assert_eq!(format!("{}", FmtN::slice(0, &v2)), "[none]".to_owned());
    assert_eq!(
        format!("{}", FmtN::slice(1, &v2)),
        "a, ... [2 total]".to_owned()
    );
    assert_eq!(format!("{}", FmtN::slice(2, &v2)), "a, b".to_owned());
    assert_eq!(format!("{}", FmtN::slice(3, &v2)), "a, b".to_owned());
    assert_eq!(format!("{}", FmtN::slice(0, &v3)), "[none]".to_owned());
    assert_eq!(
        format!("{}", FmtN::slice(1, &v3)),
        "a, ... [3 total]".to_owned()
    );
    assert_eq!(
        format!("{}", FmtN::slice(2, &v3)),
        "a, b, ... [3 total]".to_owned()
    );
    assert_eq!(format!("{}", FmtN::slice(3, &v3)), "a, b, c".to_owned());
    assert_eq!(format!("{}", FmtN::slice(4, &v3)), "a, b, c".to_owned());

    // Test FmtN::bset
    let v0: std::collections::BTreeSet<&str> = v0.into_iter().collect();
    let v1: std::collections::BTreeSet<&str> = v1.into_iter().collect();
    let v2: std::collections::BTreeSet<&str> = v2.into_iter().collect();
    let v3: std::collections::BTreeSet<&str> = v3.into_iter().collect();
    assert_eq!(format!("{}", FmtN::bset(0, &v0)), "[none]".to_owned());
    assert_eq!(format!("{}", FmtN::bset(1, &v0)), "[none]".to_owned());
    assert_eq!(format!("{}", FmtN::bset(0, &v1)), "[none]".to_owned());
    assert_eq!(format!("{}", FmtN::bset(1, &v1)), "a".to_owned());
    assert_eq!(format!("{}", FmtN::bset(2, &v1)), "a".to_owned());
    assert_eq!(format!("{}", FmtN::bset(0, &v2)), "[none]".to_owned());
    assert_eq!(
        format!("{}", FmtN::bset(1, &v2)),
        "a, ... [2 total]".to_owned()
    );
    assert_eq!(format!("{}", FmtN::bset(2, &v2)), "a, b".to_owned());
    assert_eq!(format!("{}", FmtN::bset(3, &v2)), "a, b".to_owned());
    assert_eq!(format!("{}", FmtN::bset(0, &v3)), "[none]".to_owned());
    assert_eq!(
        format!("{}", FmtN::bset(1, &v3)),
        "a, ... [3 total]".to_owned()
    );
    assert_eq!(
        format!("{}", FmtN::bset(2, &v3)),
        "a, b, ... [3 total]".to_owned()
    );
    assert_eq!(format!("{}", FmtN::bset(3, &v3)), "a, b, c".to_owned());
    assert_eq!(format!("{}", FmtN::bset(4, &v3)), "a, b, c".to_owned());

    // Test FmtN::slice debug
    use std::path::Path;
    use std::path::PathBuf;
    let a = PathBuf::from("a");
    let b = PathBuf::from("b");
    let c = PathBuf::from("c");
    let v0: Vec<&Path> = vec![];
    let v1: Vec<&Path> = vec![&a];
    let v2: Vec<&Path> = vec![&a, &b];
    let v3: Vec<&Path> = vec![&a, &b, &c];
    assert_eq!(format!("{:?}", FmtN::slice(0, &v0)), "[none]".to_owned());
    assert_eq!(format!("{:?}", FmtN::slice(1, &v0)), "[none]".to_owned());
    assert_eq!(format!("{:?}", FmtN::slice(0, &v1)), "[none]".to_owned());
    assert_eq!(format!("{:?}", FmtN::slice(1, &v1)), "\"a\"".to_owned());
    assert_eq!(format!("{:?}", FmtN::slice(2, &v1)), "\"a\"".to_owned());
    assert_eq!(format!("{:?}", FmtN::slice(0, &v2)), "[none]".to_owned());
    assert_eq!(
        format!("{:?}", FmtN::slice(1, &v2)),
        "\"a\", ... [2 total]".to_owned()
    );
    assert_eq!(
        format!("{:?}", FmtN::slice(2, &v2)),
        "\"a\", \"b\"".to_owned()
    );
    assert_eq!(
        format!("{:?}", FmtN::slice(3, &v2)),
        "\"a\", \"b\"".to_owned()
    );
    assert_eq!(format!("{:?}", FmtN::slice(0, &v3)), "[none]".to_owned());
    assert_eq!(
        format!("{:?}", FmtN::slice(1, &v3)),
        "\"a\", ... [3 total]".to_owned()
    );
    assert_eq!(
        format!("{:?}", FmtN::slice(2, &v3)),
        "\"a\", \"b\", ... [3 total]".to_owned()
    );
    assert_eq!(
        format!("{:?}", FmtN::slice(3, &v3)),
        "\"a\", \"b\", \"c\"".to_owned()
    );
    assert_eq!(
        format!("{:?}", FmtN::slice(4, &v3)),
        "\"a\", \"b\", \"c\"".to_owned()
    );
}
