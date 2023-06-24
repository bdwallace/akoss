/**
 * @file: tools_date.
 * @intro: uri编码工具类.
 * @author: zzmhot.
 * @email: zzmhot@163.com.
 * @Date: 2017/5/9 14:03.
 * @Copyright(©) 2017 by zzmhot.
 *
 */


// var dateFormat = function(cellValue, format='YY-MM-DD hh:mm:ss') {
function dateFormat(cellValue, format='YY-MM-DD hh:mm:ss') {
    var date = new Date(cellValue);
 
    var year = date.getFullYear(),
        month = date.getMonth()+1,//月份是从0开始的
        day = date.getDate(),
        hour = date.getHours(),
        min = date.getMinutes(),
        sec = date.getSeconds();
    var preArr = Array.apply(null,Array(10)).map(function(elem, index) {
        return '0'+index;
    });//开个长度为10的数组 格式为 ["00", "01", "02", "03", "04", "05", "06", "07", "08", "09"]
 
    var newTime = format.replace(/YY/g,year)
        .replace(/MM/g,preArr[month]||month)
        .replace(/DD/g,preArr[day]||day)
        .replace(/hh/g,preArr[hour]||hour)
        .replace(/mm/g,preArr[min]||min)
        .replace(/ss/g,preArr[sec]||sec);
 
    return newTime
}

export default {
    dateFormat
}