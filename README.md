A golang io.Writer that allows you to retrieve a set of what was written to it recently


To use: wrap an existing io.Writer `logger := writeLogger.NewWriter(os.Stdout)` 
Anything written to 'logger' will immediately be written to the supplied writer (in this case stdout)


Then write data as usual with `Write([]byte("myByteSlice"))`


When something has gone wrong you can use this to retrieve the debug information that was printed just beforehand:
`logger.Read(5)			//[]byte(string(Slice))
logger.ReadBuffer()		//bytes.NewBufferString("myByteSlice")
logger.ReadString()		//"myByteSlice"`



Writers are created with a default size of 10KiB. This value can be ajusted by calling `SetBufferSize(int)` before creating the writer.

Or you can create a Writer with a specific buffer size using `NewWriterWithSize(io.Writer, int)`


This was created to be able to capture useful information from the output of applications/Libraries where you might be running multiple instances concurrently logging to the same file, or where the output from stdout or stderr does not necessarily mean something has gone wrong, but the information is useful in the event of a fatal error. (ie:ffmpeg)
