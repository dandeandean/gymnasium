#include "strings.c"
#include "strings.h"

int main() {
  String *s0 = string_from("Hello");
  String *s1 = string_from("world!");
  string_print(*s0);
  string_print(*s1);
  String *s3 = string_concat(*s0, *s1);
  string_print(*s3);
  return 0;
}
