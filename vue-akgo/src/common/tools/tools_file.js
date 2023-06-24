/**
 * @file: tools_file.
 * @intro: uri编码工具类.
 * @author: zzmhot.
 * @email: zzmhot@163.com.
 * @Date: 2017/5/9 14:03.
 * @Copyright(©) 2017 by zzmhot.
 *
 */

export default new class File {
  constructor() {
  }

  //读文件
  readFile(filename){
    var fso = new ActiveXObject("Scripting.FileSystemObject");
    var f = fso.OpenTextFile(filename,1);
    var s = "";
    while (!f.AtEndOfStream)
    s += f.ReadLine()+"/n";
    f.Close();
    return s;
  }

  //写文件
  writeFile(filename, filecontent){
    var fso, f, s ;
    fso = new ActiveXObject("Scripting.FileSystemObject");
    f = fso.OpenTextFile(filename,8,true);
    f.WriteLine(filecontent);
    f.Close();
  }

}
