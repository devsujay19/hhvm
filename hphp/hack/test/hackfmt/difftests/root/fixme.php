<?hh

final class Foo {
  public function bar(
    vec<int> $ints,
  ): vec<int> {
    $ints = Vec\shuffle($ints);
    /* HH_FIXME[4110] */
    return vec[1,
               2,
               3];
  }
}
