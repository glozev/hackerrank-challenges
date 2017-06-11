#!/usr/bin/env bash

challenge=`echo $1 | tr '-' '_'`

mkdir -p src/$challenge
cat > src/$challenge/$challenge.js <<EOF
#!/usr/bin/env node
EOF
chmod +x src/$challenge/$challenge.js

mkdir -p test/$challenge
cat > test/$challenge/$challenge\_test.js << EOF
const expect = require('chai').expect;

const path = require('path');
const child = require('child_process');

const $challenge = require('../../src/$challenge/$challenge.js')

describe('$challenge', () => {

  it('should pass dummy test', () => {
    expect(true).to.be.true;
  });

  it('should read from stdin and write to stdout', (done) => {
    exec = path.join(__dirname, '../../src/$challenge/$challenge.js');
    proc = child.spawn(exec);

    proc.stdin.write('CHANGE_ME\n');
    proc.stdin.end();

    proc.stdout.once('data', (output) => {
      expect(output.toString().trim()).to.eql('CHANGE_ME');
      done();
    });
  });
});

EOF
