# Fanboi
Go fan controller primarily for unraid.  Created because although there's scripts to set fan speeds based on your disk temperatures, they're limited in what they can read / write.

features:
  plugin system for reading temperatures, writing fan speeds
  rule engine to decide when to set fan speeds based on temperatures

plugins
  liquidctl - writes fan speeds to a liquidctl config.yaml (eg: the LaaC container)
  unraiddrives - reads temperatures from the drive array in unraid

