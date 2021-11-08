import { v4 as uuidv4 } from 'uuid';

export function getGuid() {
  return uuidv4();
}

export function blurOnKey(event) {
  const { code } = event;
  if (code === 'Enter' || code === 'Escape' || code === 'Tab') {
    event.target.blur();
  }
}

export function sortOnName(array) {
  array.sort((left, right) =>
    left.name.toLowerCase().localeCompare(right.name.toLowerCase())
  );
  return array;
}
