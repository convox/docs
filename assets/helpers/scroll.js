 export default class Scroll {
  static changeOverflow(value) {
    const htmltag = document.getElementsByTagName('html')[0];
    htmltag.style.overflow = value;
  }

  static lock() {
    this.changeOverflow('hidden');
  }

  static restore() {
    this.changeOverflow('auto');
  }
}
