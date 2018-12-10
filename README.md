MacOS Proxy List
================

This project is simple wrapper library around calls to MacOS's Core Foundation framework to obtain proxy settings
from the system without needing to execute sub-processes like `networksetup` to obtain the information.

Information obtained from the Core Foundation framework is convert to native Golang types so importing projects can
make use the data returned by the package without any need to implement your own type conversions.

### Update Notice

This package will only be updated if memory leaks or security bugs are identified, no feature additions will be supported.

If you identify a memory leak or a security bug please either open an issue and a pull request or message me directly.

### Licence, Authors and Copyright

MIT License

Copyright (c) 2018 Liam Haworth <liamh@familyzone.com>

Copyright (c) 2018 Family Zone Cyber Safety Ltd

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.