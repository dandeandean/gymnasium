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
