const linkToken = artifacts.require('LinkToken')

module.exports = function (deployer) {
  deployer.deploy(linkToken)
}
