import { Controller } from "@hotwired/stimulus"

export default class extends Controller {
  connect() {
    console.log("hello")
  }

  reset() {
    console.log('reset')
    this.element.reset()
  }
}
