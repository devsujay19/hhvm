<?hh



<<__EntryPoint>>
function main_882() :mixed{
$i = 0;
 print ++$i;
 print "\t";
 print (dict['a' => 1]===true) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ===true) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = true;
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === true	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]===false) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ===false) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = false;
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === false	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]===1) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ===1) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = 1;
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === 1	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]===0) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ===0) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = 0;
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === 0	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]===-1) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ===-1) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = -1;
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === -1	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]==='1') ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ==='1') ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = '1';
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === '1'	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]==='0') ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ==='0') ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = '0';
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === '0'	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]==='-1') ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ==='-1') ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = '-1';
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === '-1'	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]===null) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ===null) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = null;
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === null	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]===dict[]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ===dict[]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict[];
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === array()	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]===vec[1]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ===vec[1]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec[1];
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === array(1)	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]===vec[2]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ===vec[2]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec[2];
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === array(2)	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]===vec['1']) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ===vec['1']) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec['1'];
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === array('1')	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]===dict['0' => '1']) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ===dict['0' => '1']) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict['0' => '1'];
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === array('0' => '1')	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]===vec['a']) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ===vec['a']) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec['a'];
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === array('a')	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]===dict['a' => 1]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ===dict['a' => 1]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict['a' => 1];
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === array('a' => 1)	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]===dict['b' => 1]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ===dict['b' => 1]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict['b' => 1];
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === array('b' => 1)	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]===dict['a' => 1, 'b' => 2]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ===dict['a' => 1, 'b' => 2]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict['a' => 1, 'b' => 2];
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === array('a' => 1, 'b' => 2)	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]===vec[dict['a' => 1]]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ===vec[dict['a' => 1]]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec[dict['a' => 1]];
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === array(array('a' => 1))	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]===vec[dict['b' => 1]]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ===vec[dict['b' => 1]]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec[dict['b' => 1]];
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === array(array('b' => 1))	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]==='php') ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ==='php') ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = 'php';
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === 'php'	";
 print "\n";
 print ++$i;
 print "\t";
 print (dict['a' => 1]==='') ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = dict['a' => 1];
 print ($a ==='') ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = '';
 print (dict['a' => 1]===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "array('a' => 1) === ''	";
 print "\n";
}
