const { expect } = require("chai");
const { ethers, network } = require("hardhat");
const {
  ContractFunctionVisibility,
} = require("hardhat/internal/hardhat-network/stack-traces/model");

describe("PrivateSale", async () => {
  let maxUnit256 =
    "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff";

  const SECONDS_IN_DAY = 86400;
  const MILISECONDS_IN_DAY = 86400000;
  const ONE_MILLION = 1000000;
  const POW18 = Math.pow(10, 18);

  before(async () => {
    [owner, addr1, addr2, addr3] = await ethers.getSigners();
    //console.log(new Date().getTime())
  });

  beforeEach(async () => {
    const IDOToken = await ethers.getContractFactory("IDOToken");
    idoToken = await IDOToken.deploy();
    await idoToken.deployed();

    currentBlock = await ethers.provider.getBlock(
      await ethers.provider.getBlockNumber()
    );
    let currentBlockTime = currentBlock.timestamp;

    let PrivateSale = await ethers.getContractFactory("PrivateSale");

    privateSale = await PrivateSale.deploy(
      idoToken.address,
      currentBlockTime + 30 * SECONDS_IN_DAY,
      currentBlockTime + 60 * SECONDS_IN_DAY,
      currentBlockTime + 90 * SECONDS_IN_DAY,
      10,
      [addr1.address, addr2.address],
      [1000, 1000]
    );
    await privateSale.deployed();

    await idoToken.connect(owner).transfer(privateSale.address, ONE_MILLION);

    await idoToken.connect(owner).approve(privateSale.address, maxUnit256); // approve owner

    //await idoToken.connect(owner).transfer(addr1.address, initSupply); // chuyển token từ owner sang addr1
    await idoToken.connect(addr1).approve(privateSale.address, maxUnit256); // approve addr1

    //await idoToken.connect(owner).transfer(addr2.address, initSupply); // chuyển token từ owner sang addr2
    await idoToken.connect(addr2).approve(privateSale.address, maxUnit256); // approve addr2
  });
  it("Should not constructor", async () => {
    let PrivateSale = await ethers.getContractFactory("PrivateSale");
    currentBlock = await ethers.provider.getBlock(
      await ethers.provider.getBlockNumber()
    );
    let currentBlockTime = currentBlock.timestamp;

    //openSaleTime < now
    await expect(
      PrivateSale.deploy(
        idoToken.address,
        currentBlockTime - 100, // XXX
        currentBlockTime + 29 * SECONDS_IN_DAY,
        currentBlockTime + 90 * SECONDS_IN_DAY,
        10,
        [addr1.address, addr2.address],
        [1000, 1000]
      )
    ).to.be.revertedWith(
      "PrivateSale/constructor: openSaleTime must after current time"
    );
    //endSaleTime < openSaleTime
    await expect(
      PrivateSale.deploy(
        idoToken.address,
        currentBlockTime + 30 * SECONDS_IN_DAY,
        currentBlockTime + 29 * SECONDS_IN_DAY, // XXX
        currentBlockTime + 90 * SECONDS_IN_DAY,
        10,
        [addr1.address, addr2.address],
        [1000, 1000]
      )
    ).to.be.revertedWith(
      "PrivateSale/constructor: endSaleTime must after openSaleTime"
    );
    //claimTime < endSaleTime
    await expect(
      PrivateSale.deploy(
        idoToken.address,
        currentBlockTime + 30 * SECONDS_IN_DAY,
        currentBlockTime + 60 * SECONDS_IN_DAY,
        currentBlockTime + 59 * SECONDS_IN_DAY, // XXX
        10,
        [addr1.address, addr2.address],
        [1000, 1000]
      )
    ).to.be.revertedWith(
      "PrivateSale/constructor: claimableTime must after endSaleTime"
    );
    // ratio == 0
    await expect(
      PrivateSale.deploy(
        idoToken.address,
        currentBlockTime + 30 * SECONDS_IN_DAY,
        currentBlockTime + 60 * SECONDS_IN_DAY,
        currentBlockTime + 90 * SECONDS_IN_DAY,
        0, // XXX
        [addr1.address, addr2.address],
        [1000, 1000]
      )
    ).to.be.revertedWith("PrivateSale/constructor: invalid ratio");
    // null array
    await expect(
      PrivateSale.deploy(
        idoToken.address,
        currentBlockTime + 30 * SECONDS_IN_DAY,
        currentBlockTime + 60 * SECONDS_IN_DAY,
        currentBlockTime + 90 * SECONDS_IN_DAY,
        10,
        [], // XXX
        [1000, 1000]
      )
    ).to.be.revertedWith("PrivateSale/constructor: null array");
    await expect(
      PrivateSale.deploy(
        idoToken.address,
        currentBlockTime + 30 * SECONDS_IN_DAY,
        currentBlockTime + 60 * SECONDS_IN_DAY,
        currentBlockTime + 90 * SECONDS_IN_DAY,
        10,
        [addr1.address, addr2.address],
        [] // XXX
      )
    ).to.be.revertedWith("PrivateSale/constructor: null array");

    // not equal array
    await expect(
      PrivateSale.deploy(
        idoToken.address,
        currentBlockTime + 30 * SECONDS_IN_DAY, // XXX
        currentBlockTime + 60 * SECONDS_IN_DAY,
        currentBlockTime + 90 * SECONDS_IN_DAY,
        10,
        [addr1.address], //XXX
        [1000, 1000]
      )
    ).to.be.revertedWith("PrivateSale/constructor: not equal array");
    // limitAmount == 0
    await expect(
      PrivateSale.deploy(
        idoToken.address,
        currentBlockTime + 30 * SECONDS_IN_DAY,
        currentBlockTime + 60 * SECONDS_IN_DAY,
        currentBlockTime + 90 * SECONDS_IN_DAY,
        10,
        [addr1.address, addr2.address],
        [1000, 0] //XXX
      )
    ).to.be.revertedWith("PrivateSale/constructor: invalid limitAmount");
    // duplicate user
    await expect(
      PrivateSale.deploy(
        idoToken.address,
        currentBlockTime + 30 * SECONDS_IN_DAY,
        currentBlockTime + 60 * SECONDS_IN_DAY,
        currentBlockTime + 90 * SECONDS_IN_DAY,
        10,
        [addr1.address, addr2.address, addr1.address], //XXX
        [1000, 1000, 300]
      )
    ).to.be.revertedWith("PrivateSale/constructor: user is already add");
  });

  it("Should buy", async () => {
    // Time + 30 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 30]);
    await ethers.provider.send("evm_mine");

    await privateSale.connect(addr1).buy({
      value: 70,
    });
    addr1Detail = await privateSale.getUserDetail(addr1.address);
    expect(addr1Detail.boughtAmount).to.equal(700);
  });
  it("Should not buy because not user", async () => {
    await expect(
      privateSale.connect(addr3).buy({
        value: 70,
      })
    ).to.be.revertedWith("PrivateSale/onlyUser: not user");
  });
  it("Should not buy because of not buy time", async () => {
    await expect(
      privateSale.connect(addr1).buy({
        value: 70,
      })
    ).to.be.revertedWith("PrivateSale/buy: not time to buy");
  });
  it("Should not buy because of max buy limit", async () => {
    // Time + 30 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 30]);
    await ethers.provider.send("evm_mine");

    await expect(
      privateSale.connect(addr1).buy({
        value: 110,
      })
    ).to.be.revertedWith("PrivateSale/buy: max buylimit");

    await privateSale.connect(addr1).buy({
      value: 70,
    });
    await expect(
      privateSale.connect(addr1).buy({
        value: 40,
      })
    ).to.be.revertedWith("PrivateSale/buy: max buylimit");
  });
  it("Should claim", async () => {
    // Time + 30 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 30]);
    await ethers.provider.send("evm_mine");

    await privateSale.connect(addr1).buy({
      value: 70,
    });
    // Time + 60 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 60]);
    await ethers.provider.send("evm_mine");
    await privateSale.connect(addr1).claim();
    addr1Detail = await privateSale.getUserDetail(addr1.address);
    expect(addr1Detail.boughtAmount).to.equal(700);
    expect(addr1Detail.claimAmount).to.equal(700);
    expect(await idoToken.balanceOf(addr1.address)).to.equal(700);
  });
  it("Should not claim because not user", async () => {
    // Time + 30 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 30]);
    await ethers.provider.send("evm_mine");

    await privateSale.connect(addr1).buy({
      value: 70,
    });
    // Time + 60 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 60]);
    await ethers.provider.send("evm_mine");
    await expect(privateSale.connect(addr3).claim()).to.be.revertedWith(
      "PrivateSale/onlyUser: not user"
    );
  });
  it("Should not claim because not buy", async () => {
    // Time + 30 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 30]);
    await ethers.provider.send("evm_mine");

    // Time + 60 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 60]);
    await ethers.provider.send("evm_mine");
    await expect(privateSale.connect(addr1).claim()).to.be.revertedWith(
      "PrivateSale/claim: not enough to claim"
    );
  });
  it("Should not claim because not time to claim", async () => {
    // Time + 30 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 30]);
    await ethers.provider.send("evm_mine");

    await privateSale.connect(addr1).buy({
      value: 70,
    });
    // Time + 30 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 30]);
    await ethers.provider.send("evm_mine");
    await expect(privateSale.connect(addr1).claim()).to.be.revertedWith(
      "PrivateSale/claim: not time to claim"
    );
  });
  it("Should not claim because not enough token in contract to claim", async () => {
    let currentBlockTime = currentBlock.timestamp;

    let PrivateSale = await ethers.getContractFactory("PrivateSale");

    privateSale = await PrivateSale.deploy(
      idoToken.address,
      currentBlockTime + 30 * SECONDS_IN_DAY,
      currentBlockTime + 60 * SECONDS_IN_DAY,
      currentBlockTime + 90 * SECONDS_IN_DAY,
      10,
      [addr1.address, addr2.address],
      [1000, 1000]
    ); //dont put let or const
    await privateSale.deployed();

    await idoToken.connect(owner).transfer(privateSale.address, 50); //XXX
    await idoToken.connect(owner).approve(privateSale.address, maxUnit256);
    await idoToken.connect(addr1).approve(privateSale.address, maxUnit256);

    // Time + 30 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 30]);
    await ethers.provider.send("evm_mine");

    await privateSale.connect(addr1).buy({
      value: 70,
    });
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 60]);
    await ethers.provider.send("evm_mine");
    await expect(privateSale.connect(addr1).claim()).to.be.revertedWith(
      "ERC20: transfer amount exceeds balance"
    );
  });
  it("Should withdraw", async () => {
    // Time + 30 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 30]);
    await ethers.provider.send("evm_mine");

    await privateSale.connect(addr1).buy({
      value: 70,
    });
    // Time + 30 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 30]);
    await ethers.provider.send("evm_mine");

    expect(await idoToken.balanceOf(privateSale.address)).to.equal(ONE_MILLION);
    expect(await ethers.provider.getBalance(privateSale.address)).to.equal(70);

    await privateSale.connect(owner).withdraw();

    expect(await idoToken.balanceOf(privateSale.address)).to.equal(700);
    expect(await ethers.provider.getBalance(privateSale.address)).to.equal(0);
    // Time + 30 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 30]);
    await ethers.provider.send("evm_mine");

    await privateSale.connect(addr1).claim();

    expect(await idoToken.balanceOf(privateSale.address)).to.equal(0);
    expect(await ethers.provider.getBalance(privateSale.address)).to.equal(0);
  });
  it("Should not withdraw because not owner", async () => {
    // Time + 30 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 30]);
    await ethers.provider.send("evm_mine");

    await privateSale.connect(addr1).buy({
      value: 70,
    });
    // Time + 30 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 30]);
    await ethers.provider.send("evm_mine");

    expect(await idoToken.balanceOf(privateSale.address)).to.equal(ONE_MILLION);
    expect(await ethers.provider.getBalance(privateSale.address)).to.equal(70);

    await expect(privateSale.connect(addr1).withdraw()).to.be.revertedWith(
      "Ownable: caller is not the owner"
    );
  });

  //TODO
  it("Should ultimate", async () => {
    //buy
    // Time + 30 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 30]);
    await ethers.provider.send("evm_mine");

    await privateSale.connect(addr1).buy({
      value: ethers.utils.parseUnits("70", "wei"),
    });

    detailAddr1 = await privateSale.getUserDetail(addr1.address);
    expect(detailAddr1.boughtAmount).to.equal(700);

    await privateSale.connect(addr1).buy({
      value: ethers.utils.parseUnits("20", "wei"),
    });
    detailAddr1 = await privateSale.getUserDetail(addr1.address);
    expect(detailAddr1.boughtAmount).to.equal(900);

    await expect(
      privateSale.connect(addr1).buy({
        value: ethers.utils.parseUnits("40", "wei"),
      })
    ).to.be.revertedWith("PrivateSale/buy: max buylimit");
    detailAddr1 = await privateSale.getUserDetail(addr1.address);
    expect(detailAddr1.boughtAmount).to.equal(900);

    //claim error
    // Time + 30 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 30]);
    await ethers.provider.send("evm_mine");
    await expect(privateSale.connect(addr1).claim()).to.be.revertedWith(
      "PrivateSale/claim: not time to claim"
    );

    //claim
    // Time + 30 days
    await ethers.provider.send("evm_increaseTime", [SECONDS_IN_DAY * 30]);
    await ethers.provider.send("evm_mine");
    await privateSale.connect(addr1).claim();
    detailAddr1 = await privateSale.getUserDetail(addr1.address);
    expect(detailAddr1.claimAmount).to.equal(900);
    expect(await idoToken.balanceOf(addr1.address)).to.equal(900);

    //withdraw
    console.log(
      "PrivateSale: ",
      await ethers.provider.getBalance(privateSale.address)
    );
    balanceBefore = await ethers.provider.getBalance(owner.address);
    console.log("before: ", balanceBefore);
    await privateSale.connect(owner).withdraw();
    balaceAfter = await ethers.provider.getBalance(owner.address);
    console.log("after: ", balaceAfter);
    console.log(balanceBefore - balaceAfter);
    //withdraw
    console.log(
      "PrivateSale: ",
      await ethers.provider.getBalance(privateSale.address)
    );
    balanceBefore = await ethers.provider.getBalance(owner.address);
    console.log("before: ", balanceBefore);
    await privateSale.connect(owner).withdraw();
    balaceAfter = await ethers.provider.getBalance(owner.address);
    console.log("after: ", balaceAfter);
    console.log(balanceBefore - balaceAfter);
  });
});
