
### Objective

Write a program called `envtool` that can be used to find all 
the files that are different between two Linux/Unix directories. 
One use case, for example, might be for seeing which files have changed 
on a machine since it was last inspected. Another use case might be for 
checking which files are different between two running docker containers.

The tool should be able to capture an environment into a summary file,
as well as compare two summary files and enumerate the differences.

The tool should keep track of all the files and directories it was
unable to process, and print a summary of those "errors" at the end.

Here is an example sequence of how the tool might be used. 
On one machine, you might run:

```
$ envtool capture --outfile /tmp/capture1 /
63002 files captured
1510 symlinks captured
could not read:
  /root/a_file
could not traverse into:
  /root/a_directory
```

Which would yield a file (`/tmp/capture1`)

Then, on a different machine, or different container, one might run

```
$ envtool capture --outfile ~/capture2 /
63000 files captured
1500 symlinks captured
could not traverse into:
  /root/a_directory
```

And by copying those two capture files to a workstation, one can use the tool
to compare them, getting this kind of example output.

```
$ envtool compare capture1 capture2
File /usr/lib64/libldap-2.4.so.2.10.7 is different 
  capture1 hash: d8330e05a90a9dad518b83611c06187cc80e6081d020bc77f3a57127e3795b88 
  capture2 hash: 1abf66644ca790507b793bb329464a97d93aa63c1911020907d754a35eb33cff
<more output from different files would follow here>
Symlink /usr/lib64/libusan.so.0 is different
  capture1 target: libusan.so.0.0.0
  capture2 target: libusan.so.0.0.1
<more output from different symlinks would follow here>
Regular file /home/user/foobar exists in capture1 but not capture2
Special file /var/run/docker.sock exists in capture2 but not capture1
```

Your output does not need to match exactly, feel free to change it as long as 
the same functionality is present.

For files that are not regular or symlinks, just compare the existence 
and type of the file

### Evaluation Criteria

The resulting program will be evaluated based on how well the code
meets the stated requirements.

In addition, the tool and code repository should be reasonably polished, 
as if this were a command line tool you were going to release. 

### Submit

Please organize, design, test and document your code and files as if it were
going into production - then push your changes to the master branch.

All the best,
Antimatter
