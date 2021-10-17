//toastr的配置信息
$(function (){
    toastr.options = {
        closeButton: true,
        debug: false,
        progressBar: false,
        positionClass: "toast-top-center",
        onclick: null,
        showDuration: "500",
        hideDuration: "500",
        timeOut: "2000",
        extendedTimeOut: "1000",
        showEasing: "swing",
        hideEasing: "linear",
        showMethod: "fadeIn",
        hideMethod: "fadeOut"
    }
})