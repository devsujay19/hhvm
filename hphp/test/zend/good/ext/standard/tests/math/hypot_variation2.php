<?hh
/* Prototype  : float hypot  ( float $x  , float $y  )
 * Description: Calculate the length of the hypotenuse of a right-angle triangle
 * Source code: ext/standard/math.c
 */

// get a class
class classA
{
}
<<__EntryPoint>> function main(): void {
echo "*** Testing hypot() : usage variations ***\n";


// heredoc string
$heredoc = <<<EOT
abc
xyz
EOT;

// get a resource variable
$fp = fopen(__FILE__, "r");

// unexpected values to be passed to $arg argument
$inputs = vec[
       // int data
/*1*/  0,
       1,
       12345,
       -2345,
       2147483647,

       // float data
/*6*/  10.5,
       -10.5,
       12.3456789000e10,
       12.3456789000E-10,
       .5,

       // null data
/*11*/ NULL,
       null,

       // boolean data
/*13*/ true,
       false,
       TRUE,
       FALSE,

       // empty data
/*17*/ "",
       '',
       vec[],

       // string data
/*20*/ "abcxyz",
       'abcxyz',
       $heredoc,

       // object data
/*23*/ new classA(),



       // resource variable
/*24*/ $fp
];

// loop through each element of $inputs to check the behaviour of hypot()
$iterator = 1;
foreach($inputs as $input) {
    echo "\n-- Iteration $iterator --\n";
    try { var_dump(hypot(3, $input)); } catch (Exception $e) { echo "\n".'Warning: '.$e->getMessage().' in '.__FILE__.' on line '.__LINE__."\n"; }
    $iterator++;
};
fclose($fp);
echo "===Done===";
}
