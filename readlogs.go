package main

import (
  "time"
  "os"
  "fmt"
  )


type LogVal struct {
   log  string
   lastModified time.Time
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}




// Tail the file for time seconds
func TailFile(log string, seconds int, start_bytes int64) string {
  time.Sleep(time.Duration(seconds) * time.Second )
  fileinfo, err1 :=  os.Stat(log)
  check (err1)
  end_bytes :=  fileinfo.Size()
  file, err2 :=  os.Open(log)
  check (err2)
  _, err3 := file.Seek(start_bytes, 0)
  check (err3)
  read_length := end_bytes - start_bytes
  bRead := make([]byte, read_length)
  bLength, err4 := file.Read(bRead)
  check (err4)
  file.Close()
  return string(bRead[:bLength])
}

// read the logs from log, for time seconds, if last != lastmodified
func ReadLogs(log string, seconds int, last time.Time) LogVal{
  file, err := os.Stat(log)
  check (err)
  bytes := file.Size()
  lastModified := file.ModTime()
  ret := ""
  // if the file hasn't changes since we last looked at it
  if lastModified != last {
    ret = TailFile(log, seconds, bytes)

  }

  return LogVal{
    log: ret,
    lastModified: lastModified,
  }
}


func main() {
  var last time.Time
  var retVal LogVal
  for  {
      retVal = ReadLogs("test.log", 5, last)
      last = retVal.lastModified
      fmt.Printf("%s", retVal.log)
      time.Sleep(time.Duration(5) * 10)

  }

}
