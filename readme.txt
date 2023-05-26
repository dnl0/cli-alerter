cli alerter telegram bot

useful when running together with processes that take too
long and you'd like to know when exactly they're finished

example:

+=======================================================+
|                                                       |
| % ./gradlew build || ./bot -m "finished compiling"    |
|                                                       |
+=======================================================+

note: you need your own bot api token and also figure out
      how to get your user id.

