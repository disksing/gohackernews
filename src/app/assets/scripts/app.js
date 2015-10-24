DOM.Ready(function() {
  // Show/Hide elements with selector in attribute data-show
  ActivateShowlinks();
  // Perform AJAX post on click on method=post|delete anchors
  ActivateMethodLinks();
});


// Perform AJAX post on click on method=post|delete anchors
function ActivateMethodLinks() {
  DOM.On('a[method="post"], a[method="delete"]', 'click', function(e) {
    // Confirm action before delete
    if (this.getAttribute('method') == 'delete') {
      if (!confirm('此操作不可撤销，是否确认删除？')) {
        return false;
      }
    }

    // Collect the authenticity token from meta tags in header
    var meta = DOM.First("meta[name='authenticity_token']")
    if (meta === undefined) {
      e.preventDefault();
      return false
    }
    var token = meta.getAttribute('content');
    
    // Perform a post to the specified url (href of link)
    var url = this.getAttribute('href');
    var redirect = this.getAttribute('data-redirect');
    var data = "authenticity_token="+token
    
    DOM.Post(url, data, function() {
      if (redirect !== null) {
        // If we have a redirect, redirect to it after the link is clicked
        window.location = redirect;
      } else {
        // If no redirect supplied, we just reload the current screen
        window.location.reload();
      }
    }, function() {
    });

    e.preventDefault();
    return false;
  });


  DOM.On('a[method="back"]', 'click', function(e) {
    history.back(); // go back one step in history
    e.preventDefault();
    return false;
  });

}


// Show/Hide elements with selector in attribute href - do this with a hidden class name
function ActivateShowlinks() {
  DOM.On('.show', 'click', function(e) {
    var selector = this.getAttribute('data-show');
    DOM.Each(selector, function(el, i) {
      if (!el.className.match(/hidden/gi)) {
        el.className = el.className + ' hidden';
      } else {
        el.className = el.className.replace(/hidden/gi, '');
      }
    });

    e.preventDefault();
    return false;
  });
}