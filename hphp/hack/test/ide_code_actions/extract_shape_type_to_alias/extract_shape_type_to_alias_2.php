<?hh

class A {
  public function foo(/*range-start*/shape('a' => int, ...)/*range-end*/ $_): void {}
}
