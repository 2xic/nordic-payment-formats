# Norwegian payment formats
There are some many "weird" banking payment file formats. Surprisingly these are the format that are the bedrock of the financial system we have here in Norway. Luckily we are starting to move to more decent file formats like ISO 20022.

### OCR 
OCR is a weird file format. Super inefficient, but hey it works. Here is the [system manual by Nets](https://www.nets.eu/no-nb/PublishingImages/Lists/Accordion%20%20OCR%20giro/AllItems/OCR%20giro%20-%20System%20manual.pdf) used for writing the parser. 

Test files used for testing was found in [python-netsgiro](https://github.com/otovo/python-netsgiro/tree/master/tests/data) and the system manual by nets.

### BGmax
BGMax is another weird file format. System manual used to implement was found [here](https://www.bankgirot.se/globalassets/dokument/tekniska-manualer/bankgiroinbetalningar_tekniskmanual_sv.pdf). The layout is more or less the same as a OCR file.