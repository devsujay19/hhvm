<?hh

function test(): void {
  type_structure(NonExistentClass::class, 'T');
}
