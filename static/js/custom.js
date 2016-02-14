/**
 * Created by LaLa on 17.01.2016.
 */

function setActiveLink() {
    var currentUrlPaths = window.location.pathname.split("/");
    var links = $('#sidebar-collapse li a');
    links.removeClass("active");
    links.each(function () {
        if ($(this).attr("href") == "/" + currentUrlPaths[1]) {
            $(this).parent().addClass("active");
        }
    });
}

$(document).ready( function() {
    setActiveLink();
});

$(function () {
    $('#hover, #striped, #condensed').click(function () {
        var classes = 'table';

        if ($('#hover').prop('checked')) {
            classes += ' table-hover';
        }
        if ($('#condensed').prop('checked')) {
            classes += ' table-condensed';
        }
        $('#table-style').bootstrapTable('destroy')
            .bootstrapTable({
                classes: classes,
                striped: $('#striped').prop('checked')
            });
    });
});