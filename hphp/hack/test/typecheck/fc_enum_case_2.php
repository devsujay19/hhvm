<?hh

enum Bar: int {
  FOO = 1;
  BAR = 2;
  BAZ = 3;
}

// Redundant match
function do_case(Bar $x): int {
  switch ($x) {
    case Bar::FOO:
      return 0;
    case Bar::BAR:
      return 1;
    case Bar::BAZ:
      return 2;
    case Bar::FOO:
      return 3;
  }
}
