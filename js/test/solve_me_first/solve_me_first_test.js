const expect = require('chai').expect;

const path = require('path');
const child = require('child_process');

const sum = require('../../src/solve_me_first/solve_me_first.js').sum

describe('solve_me_first', () => {

  it('should pass dummy test', () => {
    expect(true).to.be.true;
  });

  it('should sum two positive numbers', () => {
    expect(sum(1, 2)).to.eql(3);
  });

  it('should read from stdin and write to stdout', (done) => {
    exec = path.join(__dirname, '../../src/solve_me_first/solve_me_first.js');
    proc = child.spawn(exec);

    proc.stdin.write('1\n2\n');
    proc.stdin.end();

    proc.stdout.once('data', (output) => {
      expect(output.toString().trim()).to.eql('3');
      done();
    });
  });
});
