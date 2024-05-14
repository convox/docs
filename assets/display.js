export default class Display {
  static change(el, value) {
    el.style.display = value
  }

  static show(el, value = 'block') {
    this.change(el, value)
  }

  static hide(el) {
    this.change(el, 'none');
  }
}
