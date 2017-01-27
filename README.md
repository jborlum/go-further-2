# Deep-dive in recievers and interface types.

Over the last week I've had several discussions with people regarding receiver types, pointers, and especially the way they work with interface. 
So I took some time to investigate how they work, and I wanted to share my findings with you guys.

While learning Go I've several times encountered being unable to using pointer receivers if used in collaboration with interface. Coming from C++ this
really puzzled me. However when looking into how method expressions and method values work in Go it finally starts making sense. Unlike my
previous assumptions it is very much possible to use pointer receivers with interfaces. It just requires the correct syntax.

Hopefully you will find this useful as a reference.