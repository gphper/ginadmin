if (!WebUploader.Uploader.support()) {
    alert('上传组件不支持您的浏览器！如果你使用的是IE浏览器，请尝试升级 flash 播放器');
    throw new Error('Uploader does not support the browser you are using.');
}

var mimeArray = new Array();
mimeArray['gif'] = 'image/gif';
mimeArray['jpg'] = 'image/jpeg';
mimeArray['jpeg'] = 'image/jpeg';
mimeArray['png'] = 'image/png';
mimeArray['xml'] = 'text/xml';
mimeArray['svg'] = 'image/svg+xml';
mimeArray['xls'] = 'application/vnd.ms-excel';
mimeArray['xlsx'] = 'application/vnd.ms-excel';
mimeArray['zip'] = 'application/zip';
mimeArray['rar'] = 'application/x-rar-compressed';
mimeArray['mp3'] = 'audio/mpeg';

/**
 * @param Object o {
 *  type_key 附件类型KEY
 *  item_id 附件对应条目id
 *  pick 按钮id
 *  boxid 按钮上层div id
 *  file_path 接收file_path的input hidden id
 *  file_id 接收file_id的input hidden id
 *  callback 上传成功回调函数
 *  del_callback 删除图片回调函数
 *  extensions 文件扩展名，逗号分隔
 *  multi 是否为多图上传，true or false
 *  maximg 图片张数限制
 * }
 */
function singleUpload(o) {
    var _token = o._token,
        type_key = o.type_key,
        item_id = o.item_id,
        pick = o.pick,
        boxid = o.boxid,
        callback = o.callback,
        del_callback = o.del_callback,
        extensions = o.extensions,
        multi = o.multi,
        is_audio = o.is_audio,
        is_file = o.is_file;
    var max_size = o.maxsize;
    if (!multi) {
        multi = false;
    } else {
        multi = true;
    }
    if (!item_id) {
        item_id = 0;
    }
    var acceptOption = getAcceptOption(extensions);
    var $imgbox = $('#' + boxid);
    var sUploader = WebUploader.create({
        auto: true,
        swf: '/js/plugins/webuploader/Uploader.swf',
        server: UPLOAD_URL,
        accept: acceptOption,
        formData: {
            'ajax': 1,
            '_token': _token,
            'type_key': type_key,
            'item_id': item_id
        },
        pick: {
            id: '#' + pick,
            multiple: multi
        },
        // 是否开启压缩
        // compress: false,
        // 是否允许重复
        duplicate: true
    });
    var filecount = 0;
    sUploader.on('beforeFileQueued', function (file) {
        if (max_size && file.size > max_size) {
            layer.msg('单张照片大小不能超过' + parseInt(max_size / 1024 / 1024) + 'MB', {icon: 5, scrollbar: false, time: 2000, shade: [0.3, '#393D49']});
            return false;
        }
        if (!isNaN(o.maximg) && o.maximg > 0) {
            if ($imgbox.find(".img-div").length >= o.maximg) {
                layer.msg('图片最多上传' + o.maximg + '张', {icon: 5, scrollbar: false, time: 2000, shade: [0.3, '#393D49']});
                return false;
            }
        }
        filecount++;
    });
    var errorMsg = [];
    errorMsg['Q_TYPE_DENIED'] = '文件类型错误，请重新选择';
    errorMsg['Q_EXCEED_SIZE_LIMIT'] = '文件大小超过限制';
    errorMsg['Q_EXCEED_NUM_LIMIT'] = '文件数量超过限制';
    errorMsg['F_DUPLICATE'] = '请不要重复上传相同的文件';
    sUploader.on('error', function (code) {
        console.log(code);
        layer.msg(errorMsg[code], {icon: 5, scrollbar: false, time: 2000, shade: [0.3, '#393D49']});
    });
    sUploader.on('fileQueued', function (files) {
        $imgbox.removeClass('has-error');
        $imgbox.find('.help-block').remove();
    });
    sUploader.on('uploadError', function (file) {
        $imgbox.addClass('has-error').append('<span class="help-block"><strong>上传失败，请刷新页面后重新尝试</strong></span>');
    });
    sUploader.on('uploadComplete', function (file) {
        if (sUploader.getStats().successNum == filecount) {
            sUploader.reset();
        }
    });
    sUploader.on('uploadSuccess', function (file, response) {
        if (response.status === true) {
            // console.log(response);
            if (o.is_file) {
                var $imgdiv = $('<div id="' + response.data.file_id + '" class="img-div">' +
                    '<label>' + response.data.file_oldname + '</label>');
            } else if (o.is_audio) {
                var $imgdiv = $('<div id="' + response.data.file_id + '" class="img-div">' +
                    '<input type="text" value="' + response.data.file_path + '">');
            } else {
                var $imgdiv = $('<div id="' + response.data.file_id + '" class="img-div">' +
                    '<img src="' + response.data.file_path + '">' +
                    '<span class="cancel">×</span></div>');
                $imgdiv.on("click", "img", function () {
                    layer.photos({
                        photos: {
                            "title": "", // 相册标题
                            "id": 1, // 相册id
                            "start": 0, // 初始显示的图片序号，默认0
                            "data": [{src: $(this).attr("src"), alt: ''}]
                        },
                        shift: 5,
                        closeBtn: 2,
                        shade: 0,
                        title: false
                    });
                });
            }
            if (!multi) {
                $imgbox.find(".img-div").remove();
            }
            $imgdiv.on('click', 'span', function () {
                $(this).parent().remove();
                // 删除图片回调方法
                if (typeof del_callback === "function") {
                    del_callback(response.data);
                } else {
                    if (o.file_path) {
                        if (!multi) {
                            $('#' + o.file_path).val("");
                        } else {
                            var extfile_array = $('#' + o.file_path).val().split(',');
                            extfile_array.splice(jQuery.inArray(response.data.file_path, extfile_array), 1);
                            $('#' + o.file_path).val(extfile_array.join());
                        }
                    }
                    if (o.file_id) {
                        if (!multi) {
                            $('#' + o.file_id).val("");
                        } else {
                            var extfile_file_id_array = $('#' + o.file_id).val().split(',');
                            extfile_file_id_array.splice(jQuery.inArray(response.data.file_id + "", extfile_file_id_array), 1);
                            $('#' + o.file_id).val(extfile_file_id_array.join());
                        }
                    }
                }
            });
            $imgdiv.appendTo($imgbox);
            // 添加图片回调方法
            if (typeof callback === "function") {
                callback(response.data);
            } else {
                if (o.file_path) {
                    if (!multi) {
                        $('#' + o.file_path).val(response.data.file_path);
                    } else {
                        var file_path_val = $('#' + o.file_path).val();
                        var extfile_array = new Array();
                        if (file_path_val) {
                            extfile_array = $('#' + o.file_path).val().split(',');
                        }
                        extfile_array.push(response.data.file_path);
                        $('#' + o.file_path).val(extfile_array.join());
                    }
                }
                if (o.file_id) {
                    if (!multi) {
                        $('#' + o.file_id).val(response.data.file_id);
                    } else {
                        var file_id_val = $('#' + o.file_id).val();
                        var extfile_file_id_array = new Array();
                        if (file_id_val) {
                            extfile_file_id_array = file_id_val.split(',');
                        }
                        extfile_file_id_array.push(response.data.file_id);
                        $('#' + o.file_id).val(extfile_file_id_array.join());
                    }
                }
            }
        } else {
            $imgbox.addClass('has-error').append('<span class="help-block"><strong>' + response.msg + '</strong></span>');
        }
    });
    return sUploader;
}

function getAcceptOption(extensions) {
    var acceptOption = null;
    if (extensions !== '*') {
        if (!extensions) {
            extensions = 'gif,jpg,jpeg,png';
        }
        var mimeTypesArray = new Array();
        var extArray = extensions.split(',')
        for (var i in extArray) {
            if (mimeArray[jQuery.trim(extArray[i])]) {
                mimeTypesArray.push(mimeArray[jQuery.trim(extArray[i])]);
            }
        }
        $.unique(mimeTypesArray);
        acceptOption = {
            extensions: extensions,
            mimeTypes: mimeTypesArray.join(',')
        };
    }
    return acceptOption;
}

function sUploadDel($e, file_path_id, multi) {
    $e.parent().remove();
    if (!multi) {
        $('#' + file_path_id).val("");
    } else {
        var file_path_array = $('#' + file_path_id).val().split(',');
        file_path_array.splice(jQuery.inArray($e.parent().find('img').attr('val'), file_path_array), 1);
        $('#' + file_path_id).val(file_path_array.join());
    }
}
