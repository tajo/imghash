import fs from 'fs';
import { PNG } from 'pngjs';
import pixelmatch from 'pixelmatch';
import { compare } from "odiff-bin";

function compareImages(img1Path, img2Path) {
  const img1 = PNG.sync.read(fs.readFileSync(img1Path));
  const img2 = PNG.sync.read(fs.readFileSync(img2Path));

  const { width, height } = img1;
  const diff = new PNG({ width, height });

  const numDiffPixels = pixelmatch(
    img1.data,
    img2.data,
    diff.data,
    width,
    height,
    { threshold: 0 }
  );

  // Save diff PNG if needed
  // fs.writeFileSync('diff.png', PNG.sync.write(diff));

  return {
    totalPixels: width * height,
    differentPixels: numDiffPixels,
    percentageDiff: (numDiffPixels / (width * height)) * 100
  };
}

// Example usage
try {
  const img1 = PNG.sync.read(fs.readFileSync("fixtures/a1.png"));
  // Add this debug code
  for (let i = 0; i < 20; i += 4) {  // First 5 pixels
    console.log(`Pixel (${i / 4},0):`,
      `R:${img1.data[i]}`,
      `G:${img1.data[i + 1]}`,
      `B:${img1.data[i + 2]}`,
      `A:${img1.data[i + 3]}`);
  }
  const result = compareImages('fixtures/a1.png', 'fixtures/a2.png');
  console.log('Comparison results:', result);

  const { match, reason } = await compare(
    "fixtures/l1.png",
    "fixtures/l2.png",
    "fixtures/diffd.png",
    {
      threshold: 0.04,
      //antialiased: false,
    }
  );
  console.log(match, reason);

} catch (err) {
  console.error('Error comparing images:', err);
}

