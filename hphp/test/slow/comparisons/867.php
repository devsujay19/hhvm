<?hh



<<__EntryPoint>>
function main_867() :mixed{
$i = 0;
 print ++$i;
 print "\t";
 print (true===true) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ===true) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = true;
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === true	";
 print "\n";
 print ++$i;
 print "\t";
 print (true===false) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ===false) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = false;
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === false	";
 print "\n";
 print ++$i;
 print "\t";
 print (true===1) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ===1) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = 1;
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === 1	";
 print "\n";
 print ++$i;
 print "\t";
 print (true===0) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ===0) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = 0;
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === 0	";
 print "\n";
 print ++$i;
 print "\t";
 print (true===-1) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ===-1) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = -1;
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === -1	";
 print "\n";
 print ++$i;
 print "\t";
 print (true==='1') ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ==='1') ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = '1';
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === '1'	";
 print "\n";
 print ++$i;
 print "\t";
 print (true==='0') ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ==='0') ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = '0';
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === '0'	";
 print "\n";
 print ++$i;
 print "\t";
 print (true==='-1') ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ==='-1') ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = '-1';
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === '-1'	";
 print "\n";
 print ++$i;
 print "\t";
 print (true===null) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ===null) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = null;
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === null	";
 print "\n";
 print ++$i;
 print "\t";
 print (true===dict[]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ===dict[]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict[];
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === array()	";
 print "\n";
 print ++$i;
 print "\t";
 print (true===vec[1]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ===vec[1]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec[1];
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === array(1)	";
 print "\n";
 print ++$i;
 print "\t";
 print (true===vec[2]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ===vec[2]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec[2];
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === array(2)	";
 print "\n";
 print ++$i;
 print "\t";
 print (true===vec['1']) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ===vec['1']) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec['1'];
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === array('1')	";
 print "\n";
 print ++$i;
 print "\t";
 print (true===dict['0' => '1']) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ===dict['0' => '1']) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict['0' => '1'];
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === array('0' => '1')	";
 print "\n";
 print ++$i;
 print "\t";
 print (true===vec['a']) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ===vec['a']) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec['a'];
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === array('a')	";
 print "\n";
 print ++$i;
 print "\t";
 print (true===dict['a' => 1]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ===dict['a' => 1]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict['a' => 1];
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === array('a' => 1)	";
 print "\n";
 print ++$i;
 print "\t";
 print (true===dict['b' => 1]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ===dict['b' => 1]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict['b' => 1];
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === array('b' => 1)	";
 print "\n";
 print ++$i;
 print "\t";
 print (true===dict['a' => 1, 'b' => 2]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ===dict['a' => 1, 'b' => 2]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict['a' => 1, 'b' => 2];
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === array('a' => 1, 'b' => 2)	";
 print "\n";
 print ++$i;
 print "\t";
 print (true===vec[dict['a' => 1]]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ===vec[dict['a' => 1]]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec[dict['a' => 1]];
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === array(array('a' => 1))	";
 print "\n";
 print ++$i;
 print "\t";
 print (true===vec[dict['b' => 1]]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ===vec[dict['b' => 1]]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec[dict['b' => 1]];
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === array(array('b' => 1))	";
 print "\n";
 print ++$i;
 print "\t";
 print (true==='php') ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ==='php') ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = 'php';
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === 'php'	";
 print "\n";
 print ++$i;
 print "\t";
 print (true==='') ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = true;
 print ($a ==='') ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = '';
 print (true===$b) ? 'Y' : 'N';
 print ($a ===$b) ? 'Y' : 'N';
 print "\t";
 print "true === ''	";
 print "\n";
}
