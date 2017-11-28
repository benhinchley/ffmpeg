package ffmpeg

// FLAG                    SPEC   ARGS                 AFFECTS         IMPL
// re                      false  []                   [input]         [ ]
// accurate_seek           false  []                   [input]         [ ]
// seek_timestamp          false  []                   [input]         [ ]
// thread_queue_size       false  [size]               [input]         [ ]
// discard                 false  []                   [input]         [ ]
// tag                     true   [codec_tag]          [input output]  [ ]
// map_chapters            false  [input_file_index]   [output]        [ ]
// enc_time_base           true   [timebase]           [output]        [ ]
// bsf                     true   [bitstream_filters]  [output]        [ ]
// max_muxing_queue_size   false  [packets]            [output]        [ ]
