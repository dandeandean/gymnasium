#ifndef STRINGS_H
#define STRINGS_H

typedef struct String {
  /* Null Terminated*/
  char *chars;
  int len;
} String;

String *string_from(const char *c);
// String *string_concat(const String s1, const String s2);
char get_ith(const String *s, int i);
// String **split_at_ith(const String *s, int i);
#endif
