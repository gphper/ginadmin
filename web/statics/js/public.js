jQuery(function ($) {
    layer.config({
        title: false
    });

    $('.checkAll').on('click', function () {
        $('.checkItem').prop('checked', this.checked);
        $(this).prop('checked', this.checked);
    });

    $('.checkBtn').on('click', function () {
        if ($('.checkItem:checked').length > 0) {
            var ids = $('.checkItem:checked').map(function (index, elem) {
                return $(elem).val();
            }).get().join(',');

            btnClick($(this), ids);
        } else {
            layer.alert('请至少选择一项', {icon: 0});
        }
    });
    $(document).on('click', '.ajaxBtn', function () {
        btnClick($(this));
        return false;
    });

    var formloading = false;
    $('.ajaxForm').submit(function () {
        $(this).find('.help-block').remove();
        $(this).find('.has-error').removeClass('has-error');

        if (formloading) {
            return false;
        }

        var loading = layer.msg('系统处理中...', {icon: 16, shade: 0, time: 0});
        $.ajax({
            type: $(this).prop('method'),
            url: $(this).prop('action'),
            data: $(this).serialize(),
            dataType: 'json',
            beforeSend: function () {
                formloading = true;
            },
            success: function (data) {
                if (data.status) {
                    layer.msg(data.msg, {
                        icon: 6, scrollbar: false, time: 1000, shade: [0.3, '#393D49'], end: function () {
                            window.location.href = data.url;
                        }
                    });
                } else {
                    layer.msg(data.msg, {icon: 5, scrollbar: false, time: 2000, shade: [0.3, '#393D49']});
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                if (textStatus == 'error') {
                    let data = jqXHR.responseJSON.errors;
                    let num = 1;
                    for (let i in data) {
                        let errormsg;
                        if ($.isArray(data[i])) {
                            errormsg = data[i][0];
                        } else {
                            errormsg = data[i];
                        }

                        if (i.indexOf('.') > -1) {
                            let i_name_array = i.split('.');
                            let $ele = $('[name^=' + i_name_array[0] + ']').eq(i_name_array[1]);
                            if ($ele.parent('div').hasClass('input-group')) {
                                $ele = $ele.parent('div');
                            }
                            $ele.after('<span class="help-block"><strong>' + errormsg + '</strong></span>');
                            $ele.addClass('has-error');
                        } else {
                            let $ele = $('[name=' + i + ']');
                            if ($ele.parent('div').hasClass('input-group')) {
                                $ele = $ele.parent('div');
                            }
                            $ele.after('<span class="help-block"><strong>' + errormsg + '</strong></span>');
                            $ele.parent('div').addClass('has-error');
                            if (num == 1) {
                                $('[name=' + i + ']').focus();
                                if ($('[name=' + i + ']').parent().offset() && $('[name=' + i + ']').parent().offset().top) {
                                    $(window).scrollTop($('[name=' + i + ']').parent().offset().top - 100);
                                }
                            }
                            num++;
                        }
                    }
                }
            },
            complete: function () {
                formloading = false;
                layer.close(loading);
            }
        });
        return false;
    });

    var params = [];
    if (location.search != '') {
        params = location.search.substr(1).split('&');
    }
    var sort = '', order = '';
    for (var j = 0; j < params.length; j++) {
        var arr = params[j].split('=');
        if (arr[0] == 'sort') {
            sort = arr[1];
        }
        if (arr[0] == 'order') {
            order = arr[1];
        }
    }

    $('th[orderby]').each(function () {
        if ($(this).attr('orderby') == sort) {
            $(this).removeClass('sorting');
            if (order == 'asc') {
                $(this).addClass('sorting_asc');
            } else if (order == 'desc') {
                $(this).addClass('sorting_desc');
            }
        }
    });

    $('th[orderby]').click(function () {
        var sortName = $(this).attr('orderby');
        var found = false;
        for (var i = 0; i < params.length; i++) {
            var arr = params[i].split('=');
            if ('page' == arr[0]) {
                params[i] = 'page=1';
            } else if ('sort' == arr[0]) {
                params[i] = 'sort' + '=' + sortName;
                found = true;
            } else if ('order' == arr[0]) {
                params[i] = 'order' + '=' + (arr[1] == 'asc' ? 'desc' : 'asc');
            }
        }
        if (!found) {
            params.push('sort' + '=' + sortName);
            params.push('order=desc');
        }
        location.assign(location.protocol + "//" + location.host + location.pathname + "?" + params.join('&'));
    }).css({
        cursor: 'pointer'
    });

    $(document).on('click', '.singleOpenWindow', function () {
        var title, url;
        if (this.tagName.toLowerCase() == 'a') {
            title = $(this).text();
            url = $(this).prop('href');
        } else if (this.tagName.toLowerCase() == 'input') {
            title = $(this).val();
            url = $(this).attr('uri');
        }
        if (title && url) {
            var area_w = $(this).attr('width') || '100%';
            var area_h = $(this).attr('height') || '100%';
            layer.open({
                title: title,
                type: 2,
                area: [area_w, area_h],
                content: url
            });
        }
        return false;
    });

    // iframe中返回按钮通用点击事件
    $('#backBtn').click(function () {
        if (self != top && !$('#rIframe', parent.document.body).prop('id')) {
            parent.layer.close(parent.layer.getFrameIndex(window.name));
        } else {
            if ($(this).attr('url')) {
                location.href = $(this).attr('url');
            } else {
                window.history.back();
            }
        }
    });
    $(".img-div").on("click", "img", function () {
        var data = [{src: $(this).attr("src"), alt: ''}];
        var json = {
            "title": "", //相册标题
            "id": 1, //相册id
            "start": 0, //初始显示的图片序号，默认0
            "data": data
        };
        layer.photos({
            photos: json,
            shift: 5,
            closeBtn: 2,
            shade: 0,
            title: false
        });
    });
});

function btnClick(btn, param) {
    if (btn.attr('msg')) {
        layer.confirm(btn.attr('msg'), {icon: 0}, function (e) {
            layer.close(e);
            btnRequest(btn, param);
        }, function (e) {
            layer.close(e);
            return false;
        });
    } else {
        btnRequest(btn, param);
    }
}

function btnRequest(btn, param) {
    var url = btn.attr('uri') || btn.prop('href');
    if (param) {
        url += '/' + param;
    }

    if (btn.attr('noajax') == 1) {
        location.href = url;
    } else {
        var loading = layer.msg('系统处理中...', {icon: 16, shade: 0, time: 0});
        $.get(url, {}, function (data) {
            layer.close(loading);
            if (data.status) {
                layer.msg(data.msg, {
                    icon: 6, scrollbar: false, time: 1000, shade: [0.3, '#393D49'], end: function () {
                        if (btn.attr('callback')) {
                            eval(btn.attr('callback'));
                        } else {
                            location.href = data.url;
                        }
                    }
                });
            } else {
                layer.msg(data.msg, {icon: 2, shade: 0.3, time: 2000});
            }
        });
    }
}

var T_DOMAIN = '';
var T_PATH = '';

function thumbs(path, width, height, type) {
    var str_md5 = jQuery.md5(path);
    var fileext = path.substring(path.lastIndexOf('.') + 1);
    type = type ? type : 32;

    var thumbs_path = T_DOMAIN + T_PATH + '/thumbimg/' + str_md5[0] + '/' + str_md5[3] + '/' + width + '/' + height + '/' + type + '/' + jQuery.base64.btoa(path) + '.auto.' + fileext;

    return thumbs_path;
}
