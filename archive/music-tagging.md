# The perfect way to manage music libraries

*Compose* playlists from tags.
Tags are requirements for a song to be listed.
If no tags are specified all songs in the library will match because there are no requirements.
If you want to make every song require to be tagged as `#house` listed:

```
#house
```

If you want to listen to elevating music that is classical:

```
#elevating and #classical
```

if you want to listen to music that is either elevating trance or elevating progressive:

```
#elevating and (#trance or #progressive)
```

If you want all underground hip hop except slow underground hip hop, and dark house:

```
((#underground and #hip-hop) (not #slow)) or (#dark and #house)
```

## Tags

Emotion: `#dark` `#`
Genre:   `#trance` `#house` `#underground` `#conscious`
