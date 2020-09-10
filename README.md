# lnutnc

lnutnc is 辽宁工业大学(Liaoning University of Technology) network connector, it provides detection network and automatic login to LNUT network.

## Get started

Build the project: 

```shell
git clone https://github.com/nodece/lnutnc
cd lnutnc
make build-mipsle # Build mipsle version, It can run on pandavan system (.e.g, PHICOMM-K2P with pandavan)
make build # Build a version that runs on your own computer
```

Run the project:

```shell
lnutnc -u 20202020 -p 8888
````

Or 

```shell
LNUT_USERID=20202020 LNUT_PASSWORD=8888 lnutnc
```                                                                          
 