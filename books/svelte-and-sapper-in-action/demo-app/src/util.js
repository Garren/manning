import { v4 as uuidv4 } from 'uuid';

export const getGuid = () => uuidv4();

export function sortOnName(array) {
  array.sort((left, right) =>
    left.name.toLowerCase().localeCompare(right.name.toLowerCase())
  );
  return array;
}
