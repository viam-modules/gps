# [`MODULE` module](https://github.com/viam-modules/MODULE)

This [MODULE module](https://app.viam.com/module/viam/MODULE) implements a MODULE [MODEL COMPONENT](<LINK TO HARDWARE>), used for <DESCRIPTION> using the [`rdk:component:COMPONENT` API](https://docs.viam.com/appendix/apis/components/COMPONENT/).

> [!NOTE]
> Before configuring your COMPONENT, you must [create a machine](https://docs.viam.com/cloud/machines/#add-a-new-machine).

## Configure your MODEL COMPONENT

Navigate to the [**CONFIGURE** tab](https://docs.viam.com/configure/) of your [machine](https://docs.viam.com/fleet/machines/) in the [Viam app](https://app.viam.com/).
[Add COMPONENT / MODULE:MODEL to your machine](https://docs.viam.com/configure/#components).

On the new component panel, copy and paste the following attribute template into your COMPONENT's attributes field:

```json
{
  <ATTRIBUTES>
}
```

### Attributes

The following attributes are available for `viam:MODULE:MODEL` COMPONENTs:

<EXAMPLE !!>
| Attribute | Type | Required? | Description |
| --------- | ---- | --------- | ----------  |
| `i2c_bus` | string | **Required** | The index of the I<sup>2</sup>C bus on the board that the COMPONENT is wired to. |
| `i2c_address` | string | Optional | Default: `0x77`. The [I<sup>2</sup>C device address](https://learn.adafruit.com/i2c-addresses/overview) of the COMPONENT. |

## Example configuration

### `viam:MODULE:MODEL`
```json
  {
      "name": "<your-MODULE-MODEL-COMPONENT-name>",
      "model": "viam:MODULE:MODEL",
      "type": "COMPONENT",
      "namespace": "rdk",
      "attributes": {
      },
      "depends_on": []
  }
```

### Next Steps
- To test your COMPONENT, expand the **TEST** section of its configuration pane or go to the [**CONTROL** tab](https://docs.viam.com/fleet/control/).
- To write code against your COMPONENT, use one of the [available SDKs](https://docs.viam.com/sdks/).
- To view examples using a COMPONENT component, explore [these tutorials](https://docs.viam.com/tutorials/).