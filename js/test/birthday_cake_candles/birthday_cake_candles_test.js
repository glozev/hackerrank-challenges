const expect = require('chai').expect;

const path = require('path');
const child = require('child_process');

const birthday_cake_candles = require('../../src/birthday_cake_candles.js')

describe('birthday_cake_candles', () => {

  it('should pass dummy test', () => {
    expect(true).to.be.true;
  });

  it('should read from stdin and write to stdout', (done) => {
    exec = path.join(__dirname, '../../src/birthday_cake_candles.js');
    proc = child.spawn(exec);

    proc.stdin.write('CHANGE_ME\n');
    proc.stdin.end();

    proc.stdout.once('data', (output) => {
      expect(output.toString().trim()).to.eql('CHANGE_ME');
      done();
    });
  });
});
