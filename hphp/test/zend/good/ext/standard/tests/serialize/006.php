<?hh

class �berK��li��
{
	public $��������ber = '������';
}

<<__EntryPoint>> function main(): void {
$������ = dict['������' => '������'];

$foo = new �berK��li��();

var_dump(serialize($foo));
var_dump(unserialize(serialize($foo)));
var_dump(serialize($������));
var_dump(unserialize(serialize($������)));
}
