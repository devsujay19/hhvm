/*
   +----------------------------------------------------------------------+
   | HipHop for PHP                                                       |
   +----------------------------------------------------------------------+
   | Copyright (c) 2010-present Facebook, Inc. (http://www.facebook.com)  |
   | Copyright (c) 1997-2010 The PHP Group                                |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
*/

#include "hphp/runtime/base/type-string.h"
#include "hphp/runtime/ext/hsl/hsl_locale_libc_ops.h"

namespace HPHP {

HSLLocaleLibcOps::HSLLocaleLibcOps(
  const Locale& locale
) : m_loc(locale.get()) {
}

HSLLocaleLibcOps::~HSLLocaleLibcOps() {
}

int64_t HSLLocaleLibcOps::strlen(const String& str) const {
  return str.length();
}

String HSLLocaleLibcOps::uppercase(const String& str) const {
  return str.forEachByteFast([this](char c) { return toupper_l(c, this->m_loc); });
}

String HSLLocaleLibcOps::lowercase(const String& str) const {
  return str.forEachByteFast([this](char c) { return tolower_l(c, this->m_loc); });
}

String HSLLocaleLibcOps::foldcase(const String& str) const {
  return lowercase(str);
}

} // namespace HPHP
