const expect = require('chai').expect;

const path = require('path');
const child = require('child_process');
const src_path = '../../src/grading/grading.js'

const grade = require(src_path).grade

describe('grading', () => {
  it('should pass dummy test', () => {
    expect(true).to.be.true;
  });

  it('should grade for sample', () => {
    expect(grade([73, 67, 38, 33])).to.eql([75, 67, 40, 33]);
  });

  it('should not round when failing', () => {
    expect(grade([73, 67, 37, 33])).to.eql([75, 67, 37, 33]);
  });

  it('should always round', () => {
    expect(grade([74, 68, 38, 39])).to.eql([75, 70, 40, 40]);
  });

  it('should read from stdin and write to stdout', (done) => {
    exec = path.join(__dirname, src_path);
    proc = child.spawn(exec);

    proc.stdin.write('4\n73\n67\n38\n33\n');
    proc.stdin.end();

    proc.stdout.once('data', (output) => {
      expect(output.toString().trim()).to.eql('75\n67\n40\n33');
      done();
    });
  });
});
