#include "strings.c"

int main() {
  String *s = string_from("Hello");
  String *s1 = string_from("world!");
  printf("%s\n", s->chars);
  printf("%d\n", s->len);
  printf("%c\n", get_ith(s1,9));
  return 0;
}
