package main

import (
  "time"
  "os"
  "fmt"
  )


type LogVal struct {
   fileName     string
   log          string
   lastModified time.Time
   lastBytes    int64
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}




// Tail the file for time seconds
func TailFile(fileName string, seconds int, startBytes int64) LogVal {

  time.Sleep(time.Duration(seconds) * time.Second )
  fileinfo, err1 :=  os.Stat(fileName)
  check (err1)
  lastModified := fileinfo.ModTime()
  endBytes :=  fileinfo.Size()
  file, err2 :=  os.Open(fileName)
  check (err2)
  _, err3 := file.Seek(startBytes, 0)
  check (err3)
  readLength := endBytes - startBytes
  bRead := make([]byte, readLength)
  bLength, err4 := file.Read(bRead)
  check (err4)
  file.Close()
  return LogVal {
    fileName : fileName,
    log : string(bRead[:bLength]),
    lastModified : lastModified,
    lastBytes : endBytes,
  }
}

func ReadLogsContinously(fileName string, last time.Time, startBytes int64) LogVal {
  file, err := os.Stat(fileName)
  check (err)
  lastModified := file.ModTime()
  ret := ""
  seconds := 1
  // if the file hasn't changes since we last looked at it
  if lastModified != last {
    return TailFile(fileName, seconds, startBytes)
  }

  return LogVal{
    fileName: fileName,
    log: ret,
    lastModified: lastModified,
    lastBytes: startBytes,
  }
}

// TODO: this function needs a better name
// This is just so the user can call ReadLogsContinously and pass in the value returned by it
func ReadLogsContinouslyReusable(input LogVal) LogVal {
    return ReadLogsContinously(input.fileName, input.lastModified, input.lastBytes)
}




// read the logs from log, for time seconds, if last != lastmodified
func ReadLogs(fileName string, seconds int, last time.Time) LogVal{
  file, err := os.Stat(fileName)
  check (err)
  bytes := file.Size()
  lastModified := file.ModTime()

  // if the file hasn't changes since we last looked at it
  if lastModified != last {
    return TailFile(fileName, seconds, bytes)
  }

  return LogVal{
    fileName: fileName,
    log: "",
    lastModified: lastModified,
  }
}


// Main is just for making quick test cases, should be ignored
func main() {
  var last time.Time
  var retVal LogVal
  retVal.lastModified = last
  retVal.fileName = "test.log"
  retVal.lastBytes = 0
  for  {
      retVal = ReadLogsContinouslyReusable(retVal)
      fmt.Printf("%s", retVal.log)
      time.Sleep(time.Duration(5) * 10)

  }

}
