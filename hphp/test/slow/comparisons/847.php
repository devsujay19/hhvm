<?hh



<<__EntryPoint>>
function main_847() :mixed{
$i = 0;
 print ++$i;
 print "\t";
 print (0==true) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print ($a ==true) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = true;
 print (0==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "0 == true	";
 print "\n";
 print ++$i;
 print "\t";
 print (HH\Lib\Legacy_FIXME\eq(0, false)) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print (HH\Lib\Legacy_FIXME\eq($a, false)) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = false;
 print (HH\Lib\Legacy_FIXME\eq(0, $b)) ? 'Y' : 'N';
 print (HH\Lib\Legacy_FIXME\eq($a, $b)) ? 'Y' : 'N';
 print "\t";
 print "0 == false	";
 print "\n";
 print ++$i;
 print "\t";
 print (0==1) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print ($a ==1) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = 1;
 print (0==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "0 == 1	";
 print "\n";
 print ++$i;
 print "\t";
 print (0==0) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print ($a ==0) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = 0;
 print (0==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "0 == 0	";
 print "\n";
 print ++$i;
 print "\t";
 print (0==-1) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print ($a ==-1) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = -1;
 print (0==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "0 == -1	";
 print "\n";
 print ++$i;
 print "\t";
 print (0=='1') ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print ($a =='1') ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = '1';
 print (0==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "0 == '1'	";
 print "\n";
 print ++$i;
 print "\t";
 print (HH\Lib\Legacy_FIXME\eq(0, '0')) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print (HH\Lib\Legacy_FIXME\eq($a, '0')) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = '0';
 print (HH\Lib\Legacy_FIXME\eq(0, $b)) ? 'Y' : 'N';
 print (HH\Lib\Legacy_FIXME\eq($a, $b)) ? 'Y' : 'N';
 print "\t";
 print "0 == '0'	";
 print "\n";
 print ++$i;
 print "\t";
 print (0=='-1') ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print ($a =='-1') ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = '-1';
 print (0==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "0 == '-1'	";
 print "\n";
 print ++$i;
 print "\t";
 print (HH\Lib\Legacy_FIXME\eq(0, null)) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print (HH\Lib\Legacy_FIXME\eq($a, null)) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = null;
 print (HH\Lib\Legacy_FIXME\eq(0, $b)) ? 'Y' : 'N';
 print (HH\Lib\Legacy_FIXME\eq($a, $b)) ? 'Y' : 'N';
 print "\t";
 print "0 == null	";
 print "\n";
 print ++$i;
 print "\t";
 print (0==dict[]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print ($a ==dict[]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict[];
 print (0==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "0 == array()	";
 print "\n";
 print ++$i;
 print "\t";
 print (0==vec[1]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print ($a ==vec[1]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec[1];
 print (0==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "0 == array(1)	";
 print "\n";
 print ++$i;
 print "\t";
 print (0==vec[2]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print ($a ==vec[2]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec[2];
 print (0==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "0 == array(2)	";
 print "\n";
 print ++$i;
 print "\t";
 print (0==vec['1']) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print ($a ==vec['1']) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec['1'];
 print (0==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "0 == array('1')	";
 print "\n";
 print ++$i;
 print "\t";
 print (0==dict['0' => '1']) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print ($a ==dict['0' => '1']) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict['0' => '1'];
 print (0==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "0 == array('0' => '1')	";
 print "\n";
 print ++$i;
 print "\t";
 print (0==vec['a']) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print ($a ==vec['a']) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec['a'];
 print (0==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "0 == array('a')	";
 print "\n";
 print ++$i;
 print "\t";
 print (0==dict['a' => 1]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print ($a ==dict['a' => 1]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict['a' => 1];
 print (0==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "0 == array('a' => 1)	";
 print "\n";
 print ++$i;
 print "\t";
 print (0==dict['b' => 1]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print ($a ==dict['b' => 1]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict['b' => 1];
 print (0==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "0 == array('b' => 1)	";
 print "\n";
 print ++$i;
 print "\t";
 print (0==dict['a' => 1, 'b' => 2]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print ($a ==dict['a' => 1, 'b' => 2]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict['a' => 1, 'b' => 2];
 print (0==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "0 == array('a' => 1, 'b' => 2)	";
 print "\n";
 print ++$i;
 print "\t";
 print (0==vec[dict['a' => 1]]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print ($a ==vec[dict['a' => 1]]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec[dict['a' => 1]];
 print (0==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "0 == array(array('a' => 1))	";
 print "\n";
 print ++$i;
 print "\t";
 print (0==vec[dict['b' => 1]]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print ($a ==vec[dict['b' => 1]]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec[dict['b' => 1]];
 print (0==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "0 == array(array('b' => 1))	";
 print "\n";
 print ++$i;
 print "\t";
 print (HH\Lib\Legacy_FIXME\eq(0, 'php')) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print (HH\Lib\Legacy_FIXME\eq($a, 'php')) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = 'php';
 print (HH\Lib\Legacy_FIXME\eq(0, $b)) ? 'Y' : 'N';
 print (HH\Lib\Legacy_FIXME\eq($a, $b)) ? 'Y' : 'N';
 print "\t";
 print "0 == 'php'	";
 print "\n";
 print ++$i;
 print "\t";
 print (HH\Lib\Legacy_FIXME\eq(0, '')) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = 0;
 print (HH\Lib\Legacy_FIXME\eq($a, '')) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = '';
 print (HH\Lib\Legacy_FIXME\eq(0, $b)) ? 'Y' : 'N';
 print (HH\Lib\Legacy_FIXME\eq($a, $b)) ? 'Y' : 'N';
 print "\t";
 print "0 == ''	";
 print "\n";
}
