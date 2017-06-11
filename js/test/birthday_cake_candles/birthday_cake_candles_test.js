const expect = require('chai').expect;

const path = require('path');
const child = require('child_process');

const sub = require('../../src/birthday_cake_candles/birthday_cake_candles.js')

describe('birthday_cake_candles', () => {

  it('should pass dummy test', () => {
    expect(true).to.be.true;
  });

  it('should work for the sample', () => {
    expect(sub.birthdayCakeCandles(4, [3, 2, 1, 3])).to.eql(2);
  });

  it('should blow all', () => {
    expect(sub.birthdayCakeCandles(4, [3, 3, 3, 3])).to.eql(4);
  });

  it('should blow one', () => {
    expect(sub.birthdayCakeCandles(4, [3, 2, 2, 2])).to.eql(1);
  });

  it('should blow none', () => {
    expect(sub.birthdayCakeCandles(0, [])).to.eql(0);
  });

  it('should blow last', () => {
    expect(sub.birthdayCakeCandles(4, [4, 3, 5, 5])).to.eql(2);
  });

  it('should read from stdin and write to stdout', (done) => {
    exec = path.join(__dirname, '../../src/birthday_cake_candles/birthday_cake_candles.js');
    proc = child.spawn(exec);

    proc.stdin.write('4\n3 2 1 3\n');
    proc.stdin.end();

    proc.stdout.once('data', (output) => {
      expect(output.toString().trim()).to.eql('2');
      done();
    });
  });
});
