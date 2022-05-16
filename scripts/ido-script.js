//const hre = require("hardhat");
const { ethers, network } = require("hardhat");
async function main() {
  // idoToken.address,
  //   currentBlockTime + 30 * SECONDS_IN_DAY,
  //   currentBlockTime + 60 * SECONDS_IN_DAY,
  //   currentBlockTime + 90 * SECONDS_IN_DAY,
  //   10,
  //   [addr1.address, addr2.address],
  //   [1000, 1000];
  const SECONDS_IN_DAY = 86400;
  [owner, addr1, addr2, addr3] = await ethers.getSigners();
  const IDOToken = await ethers.getContractFactory("IDOToken");
  idoToken = await IDOToken.deploy();
  await idoToken.deployed();
  console.log("IDOToken: ", idoToken.address);
  currentBlock = await ethers.provider.getBlock(
    await ethers.provider.getBlockNumber()
  );
  let currentBlockTime = currentBlock.timestamp;

  let Vesting = await ethers.getContractFactory("Vesting");

  vesting = await Vesting.deploy(
    idoToken.address,
    currentBlockTime + 30 * SECONDS_IN_DAY,
    currentBlockTime + 60 * SECONDS_IN_DAY,
    currentBlockTime + 90 * SECONDS_IN_DAY,
    10,
    [
      "0x428863bEe015c5B0426670B8f264a89C0e55b07F",
      "0x3F8190Fb7aad9C613fF35e73b740F0f795833f9e",
    ],
    [1000, 1000]
  );
  await vesting.deployed();
  console.log("Vesting: ", vesting.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
