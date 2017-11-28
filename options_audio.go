package ffmpeg

// FLAG              SPEC   ARGS           AFFECTS         IMPL
// guess_layout_max  false  [channels]     [input]         [ ]
// ar                true   [freq]         [input output]  [ ]
// ac                true   [channels]     [input output]  [ ]
// acodec            false  [codec]        [input output]  [ ]
// aframes           false  [number]       [output]        [ ]
// aq                false  [q]            [output]        [ ]
// an                false  []             [output]        [ ]
// sample_fmt        true   [sample_fmt]   [output]        [ ]
// af                false  [filtergraph]  [output]        [ ]
// atag              false  [fourcc/tag]   [output]        [ ]
