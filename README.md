```
                  __                                        
           __    /\ \  __                                   
  ___ ___ /\_\   \_\ \/\_\    ____  __  __    ___     ___   
/' __` __`\/\ \  /'_` \/\ \  /',__\/\ \/\ \ /' _ `\  /'___\ 
/\ \/\ \/\ \ \ \/\ \L\ \ \ \/\__, `\ \ \_\ \/\ \/\ \/\ \__/ 
\ \_\ \_\ \_\ \_\ \___,_\ \_\/\____/\/`____ \ \_\ \_\ \____\
 \/_/\/_/\/_/\/_/\/__,_ /\/_/\/___/  `/___/> \/_/\/_/\/____/
                                        /\___/              
                                        \/__/               
```

# what 

a command line utility that syncs a midi device to a specific bpm

# installation

the easiest way to use this utility is to clone it, then build it locally on your machine:

```
git clone git@github.com:kellydunn/midisync
cd midisync
go build
```

# usage

once you've built the binary, you can run the utility on the command line:

```
#               device            bpm
./midisync sync /dev/snd/midiC1D1 120
```