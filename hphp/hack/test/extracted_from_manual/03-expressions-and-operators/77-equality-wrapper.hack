// @generated by hh_manual from manual/hack/03-expressions-and-operators/77-equality.md
// @codegen-command : buck run fbcode//hphp/hack/src/hh_manual:hh_manual extract fbcode/hphp/hack/manual/hack/
class MyWrapper {
  public function __construct(private mixed $m) {}
}

async function example_snippet_wrapper1(): Awaitable<void> {
  new MyWrapper(1) == new MyWrapper(1); // true
  new MyWrapper(1) == new MyWrapper(2); // false
}
