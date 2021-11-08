<script>
  // When a top-level component is invalidated, this causes svelte to update
  // any DOM elements that rely on the changed variables. Sometimes this results
  // in some properties being lost, like the selection in an input, so this
  // component not only applies a mask to a control, but it will update the
  // selection after svelte has updated the DOM.
  import { tick } from "svelte";
  export let mask;
  export let value;

  function getDigits(value) {
    // create a string containing only digits in the value.
    let newValue = "";
    for (const char of value) {
      if (char >= "0" && char <= "9") newValue += char;
    }
    return newValue;
  }

  function maskValue(digits) {
    // create a string by placing digits in the mask. if the digits are "1234567"
    // and we are using our phone number mask, this returns "(123)456-7"
    const { length } = digits;
    if (length === 0) return "";
    let value = "";
    let index = 0;
    for (const char of mask) {
      if (char === "9") {
        const nextChar = digits[index++]; // add a digit from digits
        if (!nextChar) break;
        value += nextChar;
      } else if (index < length) {
        value += char; // add a mask char
      }
    }
    return value;
  }

  // handle backspace
  // when backspace is pressed, give the browser(?) a sec to process the
  // backspace and update the dom. then get the masked value, wait until svelte
  // has updated the dom, then restore the cursor postion
  function handleKeydown(event) {
    if (event.key !== "Backspace") return;
    const { target } = event;
    let { selectionStart, selectionEnd } = target; // capture the current cursor position
    setTimeout(async () => {
      // setTimeout gives time for the backspace key to be pressed
      value = maskValue(getDigits(target.value)); // changes the value rendered by the input element
      await tick(); // wait for svelte to update the dom

      // after an update, restore the insert cursor position.
      if (selectionStart === selectionEnd) selectionStart--;
      target.setSelectionRange(selectionStart, selectionEnd);
    });
  }

  // handle printable characters like digits
  function handleKeypress(event) {
    setTimeout(async () => {
      // give time for the keypress to be processed
      const { target } = event;
      let { selectionStart } = target; // capture the current cursor position
      value = maskValue(getDigits(target.value)); // change the value rendered by the input
      await tick(); // wait for svelte to update the dom
      if (selectionStart === value.length - 1) selectionStart++; // if we're at the end of the input, move ahead one.
      const maskChar = mask[selectionStart - 1];
      if (maskChar != "9") selectionStart++; // if the mask char was inserted, move selectionStart ahead one
      target.setSelectionRange(selectionStart, selectionStart); // restore the insert cursor position
    });
  }
</script>

<p>
  When a component state has been invalidated, svelte won't immediately update
  the dom. Instead it will wait for the next microtask to see if any other
  changes need to be applied, including in other components. This avoids doing
  any unecessary work and allws the browser to batch things more effectively.
</p>
<p>
  here, input is boung to an exported 'value'. changing its value will
  invalidate the control.
</p>

<input
  maxlength={mask.length}
  on:keydown={handleKeydown}
  on:keypress={handleKeypress}
  placeholder={mask}
  bind:value
/>
