//毫秒转时分秒
export function formatDuring(mss:number) {
  // let days = parseInt(mss / (1000 * 60 * 60 * 24));
  // let hours = parseInt((mss % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
  // let minutes = parseInt((mss % (1000 * 60 * 60)) / (1000 * 60));
  // let seconds = parseInt((mss % (1000 * 60)) / 1000);
  let days = Math.trunc(mss / (1000 * 60 * 60 * 24));
  let hours = Math.trunc((mss % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
  let minutes = Math.trunc((mss % (1000 * 60 * 60)) / (1000 * 60));
  let seconds = Math.trunc((mss % (1000 * 60)) / 1000);
  let ss = mss < 1000? '<1S': ''
  let d = days? days+'D' : '';
  let h = hours? hours+'H' : '';
  let m = minutes? minutes+'M' : '';
  let s = seconds? seconds + 'S' : '';
  if(mss < 1000) {
    return ss
  } else {
    return d + ' ' + h + ' ' + m + ' ' + s
  }
};