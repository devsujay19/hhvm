<?hh <<__EntryPoint>> function main(): void {
$array = vec['a', '1', 'q6', 'h20'];

var_dump(preg_grep('/^(\d|.\d)$/', $array));
var_dump(preg_grep('/^(\d|.\d)$/', $array, PREG_GREP_INVERT));
}
