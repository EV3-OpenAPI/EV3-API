# do not import all endpoints into this module because that uses a lot of memory and stack frames
# if you need the ability to import all endpoints from this module, import them with
# from ev3api.paths.sound_tones import Api

from ev3api.paths import PathValues

path = PathValues.SOUND_TONES