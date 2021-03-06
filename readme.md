# Nordic payment formats
***No warranty is given. No complaints will be answered.***

Over the years many "weird" banking payment file formats have been created. Surprisingly these are the file format that financial system are standing on top of here in the nordics. Luckily we are starting to move to more decent file formats like ISO 20022 (camt) :)

The aim of the project was for me to play some with protobuf and grpc. In addition I wanted to take a deeper look at nordic payment formats by building a simple parser for them.

_Status: The bedrock for the parsers should be in place. There are some quirks of the Cremul format that has not been fixed. The main goal of the project has been achieved, so leaving the fixing of theses quirks as an exercise to the reader :) In addition additional tests should be added if this project should be used in any professional settings (no warranty given)_

### OCR 
OCR is a weird file format. Super inefficient, but hey it works. Here is the [system manual by Nets](https://www.nets.eu/no-nb/PublishingImages/Lists/Accordion%20%20OCR%20giro/AllItems/OCR%20giro%20-%20System%20manual.pdf) used for writing the parser. 

Test files used for writing the parser was found in [python-netsgiro](https://github.com/otovo/python-netsgiro/tree/master/tests/data) and [java-netsgiro](https://github.com/Ondkloss/java-netsgiro/tree/2fd2a2f182e5d8731ce0b37191c5c128da4217e6/src/test/resources)

### BgMax
BgMax is another weird file format. System manual used to implement was found [here](https://www.bankgirot.se/globalassets/dokument/tekniska-manualer/bankgiroinbetalningar_tekniskmanual_sv.pdf). The layout is more or less the same as a OCR file, they build upon the same ideas.

The official test files were used for testing and can be found [here](https://www.bankgirot.se/tjanster/inbetalningar/bankgiro-inbetalningar/teknisk-information/)

### Cremul
Cremul is maybe the best format if we were to compare it with BgMax and OCR.

It's more "efficient" by allow variable length, and don't use all the "filler" values like BgMax / OCR.

Used the documentation provided by [Nets](https://www.nets.eu/no-nb/SiteCollectionDocuments/Egiro/Implementation%20Guidelines%20CREMUL%20(ENG).pdf) and [Nordea](https://www.nordea.no/Images/152-69443/CREMUL-implementeringsguide.pdf) to the parser.

Test files used for testing was found in [cremul-parser](https://github.com/perspilling/cremul-parser/tree/0531ecb4a30e901e51e6f81d17a2cd764ced1f0c/test/unit/files) .
