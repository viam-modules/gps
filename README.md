# [`gps` module](https://github.com/viam-modules/gps)

This [gps module](https://app.viam.com/module/viam/gps) implements the gps [rtk-serial](#Configure-your-rtk-serial-movement_sensor), rtk-pmtk, nmea, and rtk-dual-gps movement sensors using the [`rdk:component:movement_sensor` API](https://docs.viam.com/appendix/apis/components/movement_sensor/).

A global positioning system (GPS) receives signals from satellites in the earth’s orbit to determine where it is and how fast it is going.

The `rtk-serial` and [`rtk-pmtk`](https://docs.viam.com/components/movement-sensor/gps-nmea-rtk-pmtk/) movement sensor models support [NTRIP-based](https://en.wikipedia.org/wiki/Networked_Transport_of_RTCM_via_Internet_Protocol) [real time kinematic positioning (RTK)](https://en.wikipedia.org/wiki/Real-time_kinematic_positioning) GPS units ([such as these](https://www.sparkfun.com/rtk)) and RTCM versions up to 3.3.

The chip requires a correction source to get to the required positional accuracy.
The `rtk-serial` model uses an over-the-internet correction source like an RTK reference station and sends the data over a serial connection to the [board](https://docs.viam.com/components/board/).

Follow the guide to [Set up a SparkFun RTK Reference Station](https:docs.viam.com/components/movement-sensor/set-up-base-station/) to configure a SparkFun station for use with this RTK-enabled GPS movement sensor model.

## Configure your rtk-serial movement_sensor

> [!NOTE]
> Before configuring your movement_sensor, you must [create a machine](https://docs.viam.com/cloud/machines/#add-a-new-machine).

Navigate to the [**CONFIGURE** tab](https://docs.viam.com/configure/) of your [machine](https://docs.viam.com/fleet/machines/) in the [Viam app](https://app.viam.com/).
[Add movement_sensor / gps:MODEL to your machine](https://docs.viam.com/configure/#components).

On the new component panel, copy and paste the following attribute template into your movement_sensor's attributes field:

```json
{
  "serial_path": "<path_to_serial_port>",
  "serial_baud_rate": <int>,
  "ntrip_connect_attempts": <int>,
  "ntrip_mountpoint": "<identifier>",
  "ntrip_password": "<password for NTRIP server>",
  "ntrip_url": "<URL of NTRIP server>",
  "ntrip_username": "<username for NTRIP server>"
}
```

### Attributes

The following attributes are available for `viam:gps:rtk-serial` movement_sensors:

| Attribute | Type | Required? | Description |
| --------- | ---- | --------- | ----------  |
| `serial_path`            | string | **Required** | The full filesystem path to the serial device, starting with <file>/dev/</file>. To find your serial device path, first connect the serial device to your machine, then:<ul><li>On Linux, run <code>ls /dev/serial/by-path/\*</code> to show connected serial devices, or look for your device in the output of <code>sudo dmesg \| grep tty</code>. Example: <code>"/dev/serial/by-path/usb-0:1.1:1.0"</code>.</li><li>On macOS, run <code>ls /dev/tty\* \| grep -i usb</code> to show connected USB serial devices, <code>ls /dev/tty\*</code> to browse all devices, or look for your device in the output of <code>sudo dmesg \| grep tty</code>. Example: <code>"/dev/ttyS0"</code>.</li></ul> |
| `serial_baud_rate`       | int    | Optional     | The rate at which data is sent from the sensor. <br> Default: `38400` |
| `ntrip_url`              | string | **Required** | The URL of the NTRIP server from which you get correction data. Connects to a base station (maintained by a third party) for RTK corrections. |
| `ntrip_username`         | string | Optional     | Username for the NTRIP server. |
| `ntrip_password`         | string | Optional     | Password for the NTRIP server. |
| `ntrip_connect_attempts` | int    | Optional     | How many times to attempt connection before timing out. <br> Default: `10` |
| `ntrip_mountpoint`       | string | Optional     | If you know of an RTK mountpoint near you, write its identifier here. It will be appended to NTRIP address string (for example, "nysnet.gov/rtcm/**NJMTPT1**") and that mountpoint's data will be used for corrections. |

## Example configuration

### `viam:gps:rtk-serial`
```json
  {
      "name": "<your-gps-rtk-serial-movement_sensor-name>",
      "model": "viam:gps:rtk-serial",
      "type": "movement_sensor",
      "namespace": "rdk",
      "attributes": {
        "serial_path": "/dev/serial/by-path/usb-0:1.1:1.0",
        "serial_baud_rate": 115200,
        "ntrip_connect_attempts": 12,
        "ntrip_mountpoint": "MNTPT",
        "ntrip_password": "pass",
        "ntrip_url": "http://ntrip/url",
        "ntrip_username": "usr"
      },
      "depends_on": []
  }
```

The `"serial_path"` filepath used in this example is specific to serial devices connected to Linux systems.
The `"serial_path"` filepath on a macOS system might resemble <file>"/dev/ttyUSB0"</file> or <file>"/dev/ttyS0"</file>.

> [!NOTE]
> How you connect your device to an NTRIP server varies by geographic region.
You will need to research the options available to you.
If you are not sure where to start, check out this [GPS-RTK2 Hookup Guide from SparkFun](https://learn.sparkfun.com/tutorials/gps-rtk2-hookup-guide/connecting-the-zed-f9p-to-a-correction-source).

### Next Steps
- To test your movement_sensor, expand the **TEST** section of its configuration pane or go to the [**CONTROL** tab](https://docs.viam.com/fleet/control/).
- To write code against your movement_sensor, use one of the [available SDKs](https://docs.viam.com/sdks/).
- To view examples using a movement_sensor component, explore [these tutorials](https://docs.viam.com/tutorials/).
