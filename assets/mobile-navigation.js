import Display from "./display.js";

let mobileNavbar = {};

((obj) => {
  const openSidebarButton = document.getElementById(
    "open-mobile-nav-sidebar-button"
  );
  const closeSidebarButton = document.getElementById(
    "close-mobile-nav-sidebar-button"
  );
  const mobileNavigationSidebar = document.getElementById("w-nav-overlay-0");
  const mobileNavMenu = document.getElementById("nav-menu-mobile");

  const openSidebar = () => {
    openSidebarButton.addEventListener("click", () => {
      Display.show(mobileNavigationSidebar);
      openSidebarButton.classList.remove("main-navbar-slide-block");
      closeSidebarButton.classList.add("main-navbar-slide-block");
      mobileNavigationSidebar.classList.add("main-navbar-slide-overflow");
      closeSidebarButton.classList.add("w--open");
      mobileNavMenu.setAttribute("data-nav-menu-open", "");
      mobileNavMenu.classList.remove("main-sidebar-slideout-animation");
      mobileNavMenu.classList.add("main-sidebar-slidein-animation");
    });
  };

  const closeSidebar = () => {
    closeSidebarButton.addEventListener("click", () => {
      mobileNavigationSidebar.classList.remove("main-navbar-slide-overflow");
      openSidebarButton.classList.add("main-navbar-slide-block");
      closeSidebarButton.classList.remove("main-navbar-slide-block");
      closeSidebarButton.classList.remove("w--open");
      mobileNavMenu.classList.remove("main-sidebar-slidein-animation");
      mobileNavMenu.classList.add("main-sidebar-slideout-animation");
      mobileNavMenu.removeAttribute("data-nav-menu-open");
    });
  };

  obj.initialize = () => {
    openSidebar();
    closeSidebar();
  };
})(mobileNavbar);

(() => {
  if (document.querySelector(".main-navbar")) mobileNavbar.initialize();
})();

window.addEventListener("resize", function () {
  var screenWidth = window.innerWidth;
  const openSidebarButton = document.getElementById(
    "open-mobile-nav-sidebar-button"
  );
  const closeSidebarButton = document.getElementById(
    "close-mobile-nav-sidebar-button"
  );
  const mobileNavigationSidebar = document.getElementById("w-nav-overlay-0");
  const mobileNavMenu = document.getElementById("nav-menu-mobile");

  if (screenWidth <= 991) {
    openSidebarButton.classList.add("main-navbar-slide-block");
    closeSidebarButton.classList.remove("main-navbar-slide-block");
    mobileNavMenu.classList.remove("main-navbar-slide-none");
  } else {
    openSidebarButton.classList.remove("main-navbar-slide-block");
    closeSidebarButton.classList.remove("main-navbar-slide-block");
    mobileNavMenu.classList.add("main-navbar-slide-none");
  }

  mobileNavigationSidebar.classList.remove("main-navbar-slide-overflow");
  mobileNavMenu.removeAttribute("data-nav-menu-open");
});

function adjustSidebarButtons() {
  var screenWidth = window.innerWidth;
  const openSidebarButton = document.getElementById(
    "open-mobile-nav-sidebar-button"
  );
  const closeSidebarButton = document.getElementById(
    "close-mobile-nav-sidebar-button"
  );

  if (screenWidth <= 991) {
    openSidebarButton.classList.add("main-navbar-slide-block");
    closeSidebarButton.classList.remove("main-navbar-slide-block");
  } else {
    openSidebarButton.classList.remove("main-navbar-slide-block");
    closeSidebarButton.classList.remove("main-navbar-slide-block");
  }
}

document.addEventListener("DOMContentLoaded", adjustSidebarButtons);
