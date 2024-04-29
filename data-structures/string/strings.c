#include "strings.h"
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

String *string_from(const char *c) {
  String *s = malloc(sizeof(String));
  s->chars = malloc(sizeof(char) * strlen(c));
  if (NULL == s) {
    return NULL;
  }
  strcpy(s->chars, c);
  if (NULL == s->chars) {
    return NULL;
  }
  s->len = strlen(c);
  return s;
}

char get_ith(const String *s, int i) {
  if (NULL == s || i < 0 || i > s->len) {
    return '\x00';
  }
  return s->chars[i];
}

String *string_concat(const String s1, const String s2) {
  String *s = malloc(sizeof(String));
  s->chars = malloc(sizeof(char) * (strlen(s1.chars) + strlen(s2.chars)));
  if (NULL == s) {
    return NULL;
  }
  s->len = strlen(s1.chars) + strlen(s2.chars);
  memcpy(s->chars, s1.chars, strlen(s1.chars));
  memcpy(s->chars + strlen(s1.chars), s2.chars, strlen(s2.chars));
  return s;
}

void string_print(const String s) {
  printf("String::{%s, len=%d}\n", s.chars, s.len);
}
